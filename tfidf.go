package tfidf

import (
	"bytes"
	"math"
	"strings"

	"github.com/blevesearch/go-porterstemmer"
	"github.com/lytics/multibayes"
)

var (
	// classifier is a struct with a tokenizer and a matrix
	c = multibayes.NewClassifier()
	// we'll grab the tokenizer for parsing (default ngram size is 1)
	t = c.Tokenizer
)

type TermFrequency struct {
	TermMap       map[string]int
	InverseDocMap map[string]float64
	N             int
}

func NewTermFrequencyStruct() *TermFrequency {
	return &TermFrequency{
		TermMap:       make(map[string]int),
		InverseDocMap: make(map[string]float64),
		N:             0,
	}
}

// we want the number of documents (i.e. subject line) in which a term
// occurs vs. the number of times a term appears in all documents
// to do that we will need to remove duplicate words from a subject
func (f *TermFrequency) AddDocument(document string) {
	f.N++
	// tokenize the string
	alltokens := t.Tokenize([]byte(strings.ToLower(document)))
	appearances := make(map[string]bool)
	for _, token := range alltokens {
		exclude := false
		// remove stopwords if they exist
		for _, stop := range stopbytes {
			if bytes.Equal(token.Term, stop) {
				exclude = true
				break
			}
		}
		// remove duplicates if they exist
		if exists := appearances[string(token.Term)]; !exists {
			appearances[string(token.Term)] = true
		} else {
			exclude = true
			break
		}
		if exclude {
			continue
		}

		// stem each token
		tokenString := porterstemmer.StemString(string(token.Term))

		// if the token doesn't exist, add it
		if _, ok := f.TermMap[tokenString]; !ok {
			f.TermMap[tokenString] = 1
		} else {
			// otherwise increment its count
			f.TermMap[tokenString]++
		}
	}
}

func (f *TermFrequency) InverseDocumentFrequency() {
	for term, count := range f.TermMap {
		f.InverseDocMap[term] = math.Log(float64(f.N) / float64(count))
	}
}
