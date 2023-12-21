package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	fmt.Println("WIP")
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
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
	// Verify nonce size is valid
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	// Reuse storage of 'plaintext' for ciphertext destination
	ciphertext := gcm.Seal(plaintext[:0], nonce, plaintext, nil)
	return ciphertext, nil
}
