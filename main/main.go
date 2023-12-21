package main

import (
	"fmt"
	"project-go-text-encrypter/encryption_utils"
)

func main() {
	plaintext := []byte("Hello, I am a private message - encrypt me now.")

	// Example secret key - replace with non-hardcoded 16-byte / 128 bits key
	// **Never expose your secret key**
	key := []byte("abcdefgh12345678")

	ciphertext, err := encryption_utils.Encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}

	decrypted, err := encryption_utils.Decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("- - START - -")
	fmt.Println("Pre-encryption plaintext: ", string(plaintext))
	fmt.Println("- - ENCRYPTION - -")
	fmt.Println("Ciphertext (raw): ", ciphertext)
	fmt.Printf("Ciphertext (hexadecimal): %x\n", ciphertext)
	fmt.Println("- - DECRYPTION - -")
	fmt.Println("Decrypted ciphertext (raw): ", decrypted)
	fmt.Printf("Decrypted ciphertext (string): %s\n", decrypted)
	fmt.Println("- - END - -")
}
