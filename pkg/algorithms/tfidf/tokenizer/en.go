package tokenizer

import "strings"

// EnTokenizer is English tokenizer
type EnTokenizer struct{}

// Exec implements Tokenizer interface
func (e *EnTokenizer) Exec(s string) []string {
	return strings.Fields(s)
}
