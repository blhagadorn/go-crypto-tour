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
		fmt.Println("Usage: go run main.go <text_to_decrypt>")
		os.Exit(1)
	}

	encryptedText := os.Args[1]
	key := []byte("cherrycherryapplebananagrapebana") // AES-256 requires a 32-byte key

	decrypted, err := decrypt(encryptedText, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		os.Exit(1)
	}

	fmt.Println("Decrypted:", decrypted)
}

func decrypt(encryptedText string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, key[:block.BlockSize()])
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
