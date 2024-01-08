# A Tour of Simple Crypto Operations in Golang
Articles:  
[A Tour of Go Cryptography Part 1: Hashing](https://medium.com/@bryant.hagadorn/a-tour-of-go-cryptography-part-1-hashing-421f565f02e9)

## Using the examples

### 1. Hashing
SHA256 hashing of a string

`go run 1-hashing/hash.go`

### 2. Symmetric Encryption
Encrypting and decrypting a string with AES with input from the command line

#### AES encryption of a string (input of plaintext to encrypt)
```
$ go run 2-symmetric-encryption/encrypt.go helloworld
> Encrypted: 0rP4UAf/+zPKag==
```

#### AES decryption of a string (input of ciphertext to decrypt)
```
$ go run 2-symmetric-encryption/decrypt.go 0rP4UAf/+zPKag==
> Decrypted: helloworld
```

