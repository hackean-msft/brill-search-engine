// build the vector

package ranker

import (
	"fmt"
	"math"
	"sort"

	invertedIndex "github.com/teamelehyean/brill/index"
	"github.com/teamelehyean/brill/tokenizer"
)

//Rank ranks the query and returns the document score mapping
func Rank(query string) []int64 {
	tokens := tokenizer.GetTokens(query)
	results := buildVector(tokens)

	scores := func(p1, p2 *Result) bool {
		return p1.score > p2.score
	}

	By(scores).Sort(results)
	docs := make([]int64, 0, 0)
	for _, result := range results {
		docs = append(docs, result.docID)
	}
	return docs
}

func getEntryDocuments(entries []invertedIndex.Entry) []int64 {
	docSet := make(map[int64]bool)
	docs := make([]int64, 0, 0)
	for _, entry := range entries {
		list := entry.GetPostingsList()
		for e := list.Front(); e != nil; e = e.Next() {
			posting := e.Value.(*invertedIndex.Posting)
			postingID := posting.GetDocID()
			_, ok := docSet[postingID]
			if !ok {
				docSet[postingID] = true
				docs = append(docs, postingID)
			}
		}
	}
	return docs
}

func getTokensEntries(tokens []string) []invertedIndex.Entry {
	tokensLength := len(tokens)
	entries := make([]invertedIndex.Entry, 0, tokensLength)
	for _, token := range tokens {
		entry := invertedIndex.GetEntry(token)
		entries = append(entries, entry)
	}
	return entries
}

func buildVector(tokens []string) []Result {
	results := make([]Result, 0, 0)
	collectionAmount := invertedIndex.GetCollectionAmount()
	entries := getTokensEntries(tokens)
	if len(entries) > 0 {
		docs := getEntryDocuments(entries)
		docLength := len(docs)
		scores := make([]float64, docLength, docLength)
		euclideanLength := make([]float64, docLength, docLength)
		for _, entry := range entries {
			list := entry.GetPostingsList()
			i := 0
			fmt.Println(i)
			for e := list.Front(); e != nil; e = e.Next() {
				posting := e.Value.(*invertedIndex.Posting)
				docID := posting.GetDocID()
				tf := posting.GetTermFrequency()
				for i < len(docs) {
					if docID == docs[i] {
						tfidf := TFIDF(collectionAmount, entry.GetDocFrequency(), tf)
						scores[i] += tfidf
						euclideanLength[i] += math.Pow(tfidf, 2)
						break
					}
					scores[i] += 0
					euclideanLength[i] += 0
					i++
				}
			}
		}

		val := getScores(scores, euclideanLength)
		for i, score := range val {
			results = append(results, Result{docID: docs[i], score: score})
		}
	}
	return results
}

//TFIDF Calculates the TFIDF weight given a term - document mapping
func TFIDF(N, df, tf int64) float64 {
	idf := math.Log10(float64(N) / float64(df))
	return float64(tf) * idf
}

func getScores(scrs, euclideanLengths []float64) []float64 {
	length := len(scrs)
	scores := make([]float64, length, length)
	for i := 0; i < length; i++ {
		if euclideanLengths[i] != 0 {
			scores[i] = scrs[i] / math.Sqrt(euclideanLengths[i])
		}
	}
	return scores
}

//Result represents the search result
type Result struct {
	docID int64
	score float64
}

type By func(p1, p2 *Result) bool

func (by By) Sort(results []Result) {
	ps := &resultSorter{
		results: results,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type resultSorter struct {
	results []Result
	by      func(p1, p2 *Result) bool // Closure used in the Less method.
}

func (s *resultSorter) Len() int {
	return len(s.results)
}

func (s *resultSorter) Swap(i, j int) {
	s.results[i], s.results[j] = s.results[j], s.results[i]
}

func (s *resultSorter) Less(i, j int) bool {
	return s.by(&s.results[i], &s.results[j])
}
