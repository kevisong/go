package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
)

// WordCount counts word for given strings
func WordCount(words []string) map[string]float64 {
	wordCount := map[string]float64{}
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// SHA256 SHA256 hash
func SHA256(in string) string {
	sum := sha256.Sum256([]byte(in))
	return hex.EncodeToString(sum[:])
}

func normalize(cos float64) float64 {
	return 0.5 + 0.5*cos
}

func vector(a, b map[string]float64) (vec1, vec2 []float64) {
	terms := make(map[string]interface{})
	for term := range a {
		terms[term] = nil
	}
	for term := range b {
		terms[term] = nil
	}

	for term := range terms {
		vec1 = append(vec1, a[term])
		vec2 = append(vec2, b[term])
	}

	return
}

// Cosine calculates the cosine similarity of the given vector
func Cosine(a, b map[string]float64) (sim float64) {
	vec1, vec2 := vector(a, b)

	var product, squareSumA, squareSumB float64
	for i, v := range vec1 {
		product += v * vec2[i]
		squareSumA += v * v
		squareSumB += vec2[i] * vec2[i]
	}

	if squareSumA == 0 || squareSumB == 0 {
		return 0
	}

	return normalize(product / (math.Sqrt(squareSumA) * math.Sqrt(squareSumB)))
}
