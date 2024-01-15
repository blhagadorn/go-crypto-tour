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
		fmt.Println("Usage: go run decrypt.go <encrypted_text>")
		fmt.Println("Please provide the text to decrypt")
		os.Exit(1)
	}

	encryptedText := os.Args[1]

	// Generate RSA keys (In practice, you should load the private key from a secure source)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey

	// Print the keys
	fmt.Println("Private key:", privateKey)
	fmt.Println("Public key:", publicKey)

	// Decode the base64 encoded string
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		fmt.Println("Base64 Decode Error:", err)
		os.Exit(1)
	}

	// Decrypt the text
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedBytes, nil)
	if err != nil {
		fmt.Println("Decryption error:", err)
		os.Exit(1)
	}

	fmt.Println("Decrypted:", string(decryptedBytes))
}
