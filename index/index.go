package index

import (
	"container/list"
	"fmt"
)

var invertedIndex = make(map[string]*Entry)
var documentsSet = make(map[int64]bool)
var amount int64

//Term represents and inverted index term
type Term string

//Entry an inverte index entry
type Entry struct {
	docFrequency int64
	postingsList *list.List
}

// NewEntry Creates a new inverted index entry
func NewEntry(posting *Posting) *Entry {
	postingsList := list.New()
	postingsList.PushBack(posting)
	return &Entry{docFrequency: 1, postingsList: postingsList}
}

// GetPostingsList Retrieves the postings list for a particular inverted index term
func (entry *Entry) GetPostingsList() list.List {
	return *entry.postingsList
}

// GetPostingsList Retrieves the postings list for a particular inverted index term
func (entry *Entry) getPostingsList() *list.List {
	return entry.postingsList
}

// GetDocFrequency returns the document frequency for the entry
func (entry *Entry) GetDocFrequency() int64 {
	return entry.docFrequency
}

//AddPosting adds posting to the postings list for this entry
func (entry *Entry) AddPosting(p *Posting) {
	postingsList := entry.getPostingsList()
	added := false
	for e := postingsList.Front(); e != nil; e = e.Next() {
		posting := e.Value.(*Posting)
		if posting.docID > p.docID {
			postingsList.InsertBefore(p, e)
			entry.docFrequency++
			added = true
			break
		} else if posting.docID == p.docID {
			posting.tf++
			added = true
			break
		}
	}
	if !added {
		postingsList.PushBack(p)
		entry.docFrequency++
	}
}

// AddToIndex Adds a term to the inverted index
func AddToIndex(term string, docID int64) {
	_, ok := documentsSet[docID]
	if !ok {
		documentsSet[docID] = true
		incrementAmount(1)
	}
	documentsSet[docID] = true
	entry, ok := invertedIndex[term]
	if ok {
		p := NewPosting(docID, 1)
		entry.AddPosting(p)
	} else {
		posting := NewPosting(docID, 1)
		entry := NewEntry(posting)
		invertedIndex[term] = entry
	}
}

// // TFIDF calculates the tf idf score for each term in the entry
// func TFIDF(entry *Entry) {
// 	df := entry.docFrequency
// 	postingsList := entry.GetPostingsList()

// 	for e := postingsList.Front(); e != nil; e = e.Next() {
// 		posting := e.Value.(*Posting)
// 	}

// }

func DisplayInvertedIndex() {
	for key, entry := range invertedIndex {
		fmt.Printf("Key: %40s", key)
		postingsList := entry.GetPostingsList()
		for e := postingsList.Front(); e != nil; e = e.Next() {
			posting := e.Value.(*Posting)
			fmt.Printf("\tPosting: %v", posting)
		}
		fmt.Printf("\n")
	}
}

// GetEntry gets the entry for a index term
func GetEntry(term string) Entry {
	entry, ok := invertedIndex[term]
	if ok {
		return *entry
	}

	entry = &Entry{docFrequency: 0, postingsList: list.New()} // creates empty entry
	return *entry
}

func incrementAmount(value int64) {
	amount += value
}

//GetCollectionAmount returns the amount of documents in the collection
func GetCollectionAmount() int64 {
	return amount
}
