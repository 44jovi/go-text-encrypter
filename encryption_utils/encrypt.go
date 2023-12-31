package encryption_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// Create block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create GCM cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Random nonce for each encryption
	nonce := make([]byte, gcm.NonceSize())
	// Fill the nonce byte slice with random bytes
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	// TODO: review best practices / alternatives to using nonce for "dst"
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}
