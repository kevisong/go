package encoding

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// Base64Encode encodes bytes to base64 string
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// Base64Decode decodes base64 string to bytes
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// UnicodeDecode decodes unicode string to bytes
func UnicodeDecode(str string) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
