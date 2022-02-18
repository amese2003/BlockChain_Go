package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	testblock := block{"test block", "", ""}
	hash := sha256.Sum256([]byte(testblock.data + testblock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	testblock.hash = hexHash
	fmt.Println(testblock)
}
