// hasher.go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "Hello, World!"
	sum := sha256.Sum256([]byte(data))
	fmt.Printf("SHA-256 Hash: %x\n", sum)
}
