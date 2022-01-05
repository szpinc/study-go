package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	log.Println(s)
}
