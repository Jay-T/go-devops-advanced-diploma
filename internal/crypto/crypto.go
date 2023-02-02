package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type CryptoService struct {
	key []byte
}

func NewCryptoService(keyString string) *CryptoService {
	key := []byte("the-key-has-to-be-32-bytes-long!")
	return &CryptoService{key}
}

func (cs *CryptoService) Encrypt(plaintext string) (string, error) {
	plainBytes := []byte(plaintext)

	c, err := aes.NewCipher(cs.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipher := gcm.Seal(nonce, nonce, plainBytes, nil)
	return base64.StdEncoding.EncodeToString(cipher), nil
}

func (cs *CryptoService) Decrypt(ciphertext string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(cs.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherBytes) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherBytes := cipherBytes[:nonceSize], cipherBytes[nonceSize:]
	decipherBytes, err := gcm.Open(nil, nonce, cipherBytes, nil)
	if err != nil {
		return "", err
	}

	return string(decipherBytes), nil
}
