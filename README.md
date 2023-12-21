# Text Encrypter / Decrypter (Go)

A plaintext encrypter, using a symmetrical block cipher within the built-in `crypto` library in Go.

## Getting started

- Install Go on your local system
- Clone this repository
- Navigate to the repository root directory
- Use Go to run `main.go` in the `main` package

  ```
  go run ./main/main.go
  ```

- Example output:

  ```
  - - START - -

  - - PRE-ENCRYPTION - -

  Pre-encryption text (plaintext):  Hello, I am a private message - encrypt me now.

  - - ENCRYPTION - -

  Ciphertext (raw):  [171 147 191 236 206 3 197 106 13 176 39 104 77 99 17 224 39 248 111 49 153 66 24 78 39 68 79 249 94 13 248 130 31 5 1 121 48 81 80 244 173 26 191 181 124 148 48 165 184 80 104 115 145 151 135 72 63 33 145 136 130 91 50 204 167 28 15 167 224 223 142 196 119 36 170]

  Ciphertext (hexadecimal): ab93bfecce03c56a0db027684d6311e027f86f319942184e27444ff95e0df8821f050179305150f4ad1abfb57c9430a5b8506873919787483f219188825b32cca71c0fa7e0df8ec47724aa

  - - DECRYPTION - -

  Decrypted ciphertext (raw):  [72 101 108 108 111 44 32 73 32 97 109 32 97 32 112 114 105 118 97 116 101 32 109 101 115 115 97 103 101 32 45 32 101 110 99 114 121 112 116 32 109 101 32 110 111 119 46]

  Decrypted ciphertext (plaintext): Hello, I am a private message - encrypt me now.

  - - END - -
  ```

## Going further

- Try editing the pre-encrypted text - `plaintext`
- Try using a different secret - `key`
  - Ensure it is 16-bytes / 128 bits

## Disclaimer / reminder

- This encrypter/decrypter is for **general demonstration puroses only**.
- **Never wrongly expose real secrets / secret keys**.
