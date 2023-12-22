package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <text_to_encrypt>")
		os.Exit(1)
	}

	text := os.Args[1]
	key := []byte("cherrycherryapplebananagrapebana") // AES-256 requires a 32-byte key

	encrypted, err := encrypt(text, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		os.Exit(1)
	}

	fmt.Println("Encrypted:", encrypted)
}

func encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
