package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str1 := "test"
	str2 := "test"

	hash1 := sha256.Sum256([]byte(str1))
	hash2 := sha256.Sum256([]byte(str2))

	fmt.Printf("hash256(%s):%x\n", str1, hash1)
	fmt.Printf("hash256(%s):%x\n", str2, hash2)

	count := 0

	for index, cha1 := range hash1 {
		cha2 := hash2[index]
		if cha1 != cha2 {
			count++
		}
	}
	fmt.Printf("不同字符的个数:%d\n", count)
}
