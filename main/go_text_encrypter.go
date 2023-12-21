package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	plaintext := []byte("Hello, I am a private message - encrypt me now.")

	// Example secret key - replace with non-hardcoded 16-byte / 128 bits key
	// **Never expose your secret key**
	key := []byte("abcdefgh12345678")

	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("- - START - -")
	fmt.Println("Plaintext: ", string(plaintext))
	fmt.Println("- - ENCRYPTION - -")
	fmt.Println("Ciphertext (raw): ", ciphertext)
	fmt.Printf("Ciphertext (hexadecimal): %x", ciphertext)
	fmt.Println("\n- - END - -")
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
