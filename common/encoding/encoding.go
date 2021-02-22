package encoding

import (
	"crypto/rand"
	"encoding/base64"
)

// GetRandomString generate random string by specify chars.
// source: https://github.com/gogits/gogs/blob/9ee80e3e5426821f03a4e99fad34418f5c736413/modules/base/tool.go#L58
func GetRandomString(n int, alphabets ...byte) (string, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes), nil
}

// Base64Encode encodes bytes to base64 string
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// Base64Decode decodes base64 string to bytes
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
