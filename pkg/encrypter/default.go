package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type DefaultEncrypter struct{}

var _ Encrypter = (*DefaultEncrypter)(nil)

func NewDefaultEncrypter() Encrypter {
	return &DefaultEncrypter{}
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func (e *DefaultEncrypter) Encrypt(secret, raw string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	plainText := []byte(raw)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (e *DefaultEncrypter) Decrypt(secret, encrypted string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
