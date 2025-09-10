package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {

	// SHA-256
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte("Hello, Radosław!"))
	sha256Sum := sha256Hasher.Sum(nil)
	fmt.Printf("SHA-256: %x\n", sha256Sum)

	// SHA-512
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte("Hello, Radosław!"))
	sha512Sum := sha512Hasher.Sum(nil)
	fmt.Printf("SHA-512: %x\n", sha512Sum)

	// SHA-512/256
	sha512_256Hasher := sha512.New512_256()
	sha512_256Hasher.Write([]byte("Hello, Radosław!"))
	sha512_256Sum := sha512_256Hasher.Sum(nil)
	fmt.Printf("SHA-512/256: %x\n", sha512_256Sum)

	// Streaming Hash with hash.Hash interface
	var h hash.Hash = sha256.New()
	h.Write([]byte("Hello "))
	h.Write([]byte("Radosław!"))
	sum := h.Sum(nil)
	fmt.Printf("Streaming SHA-256: %x\n", sum)

	// Hashing Files (SHA-256 / SHA-512)
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	fileHasher := sha256.New()
	if _, err := io.Copy(fileHasher, file); err != nil {
		fmt.Println("File read error:", err)
		return
	}
	fileSum := fileHasher.Sum(nil)
	fmt.Printf("File SHA-256: %x\n", fileSum)
}
