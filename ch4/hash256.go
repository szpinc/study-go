package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	str := os.Args[1]

	chars := sha256.Sum256([]byte(str))

	fmt.Printf("%x\n", chars)
}
