package index

// Posting represents an inverted index posting
type Posting struct {
	docID int
	tf    int
}

// NewPosting Creates a new posting
func NewPosting(docID, tf int) *Posting {
	return &Posting{docID: docID, tf: tf}
}
