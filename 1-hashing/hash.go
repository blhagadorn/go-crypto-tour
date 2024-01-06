package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: hasher <text_to_hash>")
		fmt.Println("Please provide the text to hash")
		os.Exit(1)
	}

	// The first argument is always the program name,
	// so the actual input will be the second argument.
	data := os.Args[1]

	// Compute the SHA-256 hash
	sum := sha256.Sum256([]byte(data))
	fmt.Printf("SHA-256 Hash: %x\n of text: %s", sum, data)
}
