package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"strings"
)

// MD5Sum calculates the MD5 of a stream
func MD5Sum(reader io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, reader)
	if err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}

// MD5SumSumString calculates the MD5 of a string
func MD5SumSumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return MD5Sum(buffer)
}

// SHA256 calculates the SHA256 of a stream
func SHA256(reader io.Reader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// SHA256SumString calculates the MD5 of a string
func SHA256SumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return SHA256(buffer)
}
