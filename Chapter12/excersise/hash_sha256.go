package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func hashSHA256() {
	hasher := sha256.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v\n", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Resutl: %x\n", checksum)
}
