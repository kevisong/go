package encoding

import (
	"encoding/base64"
)

// Base64Encode encodes bytes to base64 string
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// Base64Decode decodes base64 string to bytes
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
