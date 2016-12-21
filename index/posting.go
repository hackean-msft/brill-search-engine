package index

// Posting represents an inverted index posting
type Posting struct {
	docID int64
	tf    int64
}

// NewPosting Creates a new posting
func NewPosting(docID, tf int64) *Posting {
	return &Posting{docID: docID, tf: tf}
}

// GetDocID Returns the doc id for this posting
func (p *Posting) GetDocID() int64 {
	return p.docID
}

// GetTermFrequency returns the term frequency for this posting
func (p *Posting) GetTermFrequency() int64 {
	return p.tf
}
