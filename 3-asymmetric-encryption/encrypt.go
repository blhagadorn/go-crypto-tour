package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run encrypt.go <text_to_encrypt>")
		fmt.Println("Please provide the text to encrypt")
		os.Exit(1)
	}

	text := os.Args[1]

	// Generate RSA keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey

	// Print the keys
	fmt.Println("Private key:", privateKey)
	fmt.Println("Public key:", publicKey)

	// Encrypt the text
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(text), nil)
	if err != nil {
		fmt.Println("Encryption error:", err)
		os.Exit(1)
	}
	encryptedText := base64.StdEncoding.EncodeToString(encryptedBytes)

	fmt.Println("Encrypted:", encryptedText)
}
