package index

import (
	"container/list"
	"fmt"
)

var invertedIndex = make(map[string]*Entry)

//Term represents and inverted index term
type Term string

//Entry an inverte index entry
type Entry struct {
	docFrequency int
	postingsList *list.List
}

// NewEntry Creates a new inverted index entry
func NewEntry(posting *Posting) *Entry {
	postingsList := list.New()
	postingsList.PushBack(posting)
	return &Entry{docFrequency: 1, postingsList: postingsList}
}

// GetPostingsList Retrieves the postings list for a particular inverted index term
func (e *Entry) GetPostingsList() *list.List {
	return e.postingsList
}

//AddPosting adds posting to the postings list for this entry
func (e *Entry) AddPosting(p *Posting) {
	postingsList := e.GetPostingsList()
	added := false
	for e := postingsList.Front(); e != nil; e = e.Next() {
		posting := e.Value.(*Posting)
		if posting.docID > p.docID {
			postingsList.InsertBefore(p, e)
			added = true
			break
		}
	}
	if !added {
		postingsList.PushFront(p)
	}
}

// AddToIndex Adds a term to the inverted index
func AddToIndex(term string, docID int, tf int) {
	entry, ok := invertedIndex[term]
	if ok {
		p := NewPosting(docID, tf)
		entry.AddPosting(p)
		postingsList := entry.GetPostingsList()
		fmt.Println(postingsList.Len())
		entry.docFrequency++
	} else {
		posting := NewPosting(docID, tf)
		entry := NewEntry(posting)
		invertedIndex[term] = entry
	}
}

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
