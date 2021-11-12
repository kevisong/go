package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// MD5Sum calculates the MD5 of a string
func MD5Sum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SHA256Sum calculates the SHA256Sum of a string
func SHA256Sum(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
