package tokenizer

// Tokenizer interface
type Tokenizer interface {
	Exec(string) []string
}
