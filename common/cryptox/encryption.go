package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const saltLength = 8

// Decrypt decrypts a payload with a given secret.
func Decrypt(payload []byte, secret string) ([]byte, error) {
	salt := payload[:saltLength]
	key, err := encryptionKeyToBytes(secret, string(salt))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(payload) < aes.BlockSize {
		return nil, errors.New("payload too short")
	}
	iv := payload[saltLength : saltLength+aes.BlockSize]
	payload = payload[saltLength+aes.BlockSize:]
	payloadDst := make([]byte, len(payload))

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(payloadDst, payload)
	return payloadDst, nil
}

// Encrypt encrypts a payload with a given secret.
func Encrypt(payload []byte, secret string) ([]byte, error) {
	salt, err := GetRandomString(saltLength)
	if err != nil {
		return nil, err
	}

	key, err := encryptionKeyToBytes(secret, salt)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, saltLength+aes.BlockSize+len(payload))
	copy(ciphertext[:saltLength], salt)
	iv := ciphertext[saltLength : saltLength+aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[saltLength+aes.BlockSize:], payload)

	return ciphertext, nil
}

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

// EncodePassword encodes a password using PBKDF2.
func EncodePassword(password string, salt string) (string, error) {
	newPasswd := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return hex.EncodeToString(newPasswd), nil
}

// Key needs to be 32bytes
func encryptionKeyToBytes(secret, salt string) ([]byte, error) {
	return pbkdf2.Key([]byte(secret), []byte(salt), 10000, 32, sha256.New), nil
}
