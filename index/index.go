package index

import "container/list"

var invertedIndex = make(map[string]Entry)

//Term represents and inverted index term
type Term string

//Entry an inverte index entry
type Entry struct {
	term         Term
	docFrequency int32
	postingsList list.List
}

// Posting represents an inverted index posting
type Posting struct {
	docID int32
	tf    int32
}

// NewPosting Creates a new posting
func NewPosting(docID, tf int32) *Posting {
	return &Posting{docID: docID, tf: tf}
}
