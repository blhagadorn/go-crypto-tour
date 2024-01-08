package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run hmac.go key your-message")
		fmt.Println("Please provide the text to hash")
		os.Exit(1)
	}
	// Secret key
	secretKey := []byte(os.Args[1])

	// Message to be hashed
	message := []byte(os.Args[2])

	// Create a new HMAC using SHA256
	hmacSHA256 := hmac.New(sha256.New, secretKey)

	// Write data to it
	hmacSHA256.Write(message)

	// Get the HMAC result
	hmacResult := hmacSHA256.Sum(nil)

	// Encode to hex string
	fmt.Println("HMAC (SHA256):", hex.EncodeToString(hmacResult))
}
