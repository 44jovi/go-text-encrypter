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
	fmt.Println("\n- - PRE-ENCRYPTION - -")
	fmt.Println("\nPre-encryption text (plaintext): ", string(plaintext))
	fmt.Println("\n- - ENCRYPTION - -")
	fmt.Println("\nCiphertext (raw): ", ciphertext)
	fmt.Printf("\nCiphertext (hexadecimal): %x\n", ciphertext)
	fmt.Println("\n- - DECRYPTION - -")
	fmt.Println("\nDecrypted ciphertext (raw): ", decrypted)
	fmt.Printf("\nDecrypted ciphertext (plaintext): %s\n", decrypted)
	fmt.Println("\n- - END - -")
}
