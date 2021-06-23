package tfidf

import (
	"math"
	"sync"

	"github.com/KEVISONG/go/common/algorithms/tfidf/tokenizer"
	"github.com/KEVISONG/go/common/algorithms/tfidf/utils"
)

// TFIDF defines TFIDF
type TFIDF struct {
	Corpus *Corpus

	t  tokenizer.Tokenizer
	mu *sync.RWMutex
}

// Corpus defines Corpus
type Corpus struct {
	// TermCount stores all terms appeared in the corpus as key
	// and the # of docs containing the term as value. It is
	// used for calculating idf.
	TermCount map[string]float64
	// Corpus stores all documents with it's content hash as key
	// and the document as value. It is used for calculating idf.
	Documents map[string]Document
}

// Document defines Document
type Document struct {
	ID      string
	Content string
	Terms   []string
	tfidfs  map[string]float64
}

// NewTFIDF factory
func NewTFIDF() *TFIDF {
	return &TFIDF{
		Corpus: &Corpus{
			TermCount: map[string]float64{},
			Documents: map[string]Document{},
		},
		t:  &tokenizer.EnTokenizer{},
		mu: &sync.RWMutex{},
	}
}

// NewTFIDFWithTokenizer factory
func NewTFIDFWithTokenizer(t tokenizer.Tokenizer) *TFIDF {
	return &TFIDF{
		Corpus: &Corpus{
			TermCount: map[string]float64{},
			Documents: map[string]Document{},
		},
		t:  t,
		mu: &sync.RWMutex{},
	}
}

func (t *TFIDF) calTF(terms []string) map[string]float64 {
	termCount := utils.WordCount(terms)
	tfs := map[string]float64{}
	for _, term := range terms {
		tfs[term] = termCount[term] / float64(len(terms))
	}
	return tfs
}

func (t *TFIDF) calIDF(terms []string) map[string]float64 {
	idfs := map[string]float64{}
	for _, term := range terms {
		var denominator float64
		denominator, ok := t.Corpus.TermCount[term]
		if !ok {
			denominator = 0
		}
		idfs[term] = float64(
			math.Log(
				float64(
					//(float64(len(t.Corpus)) + 1) / float64(int(denominator)+1),
					(float64(len(t.Corpus.Documents))) / float64(int(denominator)+1),
				),
			),
		)
	}
	return idfs
}

// AddDoc adds doc to the corpus by:
// 1. Update Corupus for later calculation of other doc's idf (as numerator)
// 2. Update TermDocMap for later calculation of other docs's idf (as denominator)
func (t *TFIDF) AddDoc(doc string) {

	// Update Corpus
	id := utils.SHA256(doc)
	terms := t.t.Exec(doc)
	if len(terms) == 0 {
		return
	}

	t.Corpus.Documents[id] = Document{
		ID:      id,
		Content: doc,
		Terms:   terms,
	}

	// Update TermDocMap
	for _, term := range terms {
		t.Corpus.TermCount[term]++
	}

	// Recalculate tfidf
	t.CalAll()

}

// Add Docs adds doc in batch
func (t *TFIDF) AddDocs(docs ...string) {

	for _, doc := range docs {
		t.AddDoc(doc)
	}

	// Recalculate tfidf
	t.CalAll()

}

// CalAll calculates tf-idf for all documents in the corpus
func (t *TFIDF) CalAll() {
	for id, doc := range t.Corpus.Documents {
		for term, tfidf := range t.cal(doc.Terms) {
			if doc.tfidfs == nil {
				doc.tfidfs = map[string]float64{}
			}
			doc.tfidfs[term] = tfidf
		}
		t.Corpus.Documents[id] = doc
	}
}

// Cal calculates tfidf for the doc
func (t *TFIDF) Cal(doc string) map[string]float64 {

	terms := t.t.Exec(doc)

	return t.cal(terms)

}

// cal calculates tfidf for given terms by:
// 1. Calculate tf of the doc
// 2. Calculate idf of the doc
// 2. Calculate tf-idf of the doc by tf * idf
func (t *TFIDF) cal(terms []string) map[string]float64 {

	tfidfs := map[string]float64{}

	// Calculate TF for each word in the doc
	tfs := t.calTF(terms)

	// Calculate IDF for each word in the corpus
	idfs := t.calIDF(terms)

	// Calculate TF-IDF for each word
	for term, tf := range tfs {
		tfidfs[term] = tf * idfs[term]
	}

	return tfidfs

}

// Query returns the calculated similarities of all
// the document in the corpus with the given doc.
func (t *TFIDF) Query(doc string) map[string]float64 {

	sims := map[string]float64{}

	tfidfs := t.cal(t.t.Exec(doc))

	for id, document := range t.Corpus.Documents {
		sim := utils.Cosine(tfidfs, document.tfidfs)
		sims[id] = sim
	}

	return sims

}
