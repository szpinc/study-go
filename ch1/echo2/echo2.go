package main

import (
	"log"
	"os"
)

func main() {

	for _, arg := range os.Args {
		log.Println(arg)
	}

	for index, arg := range os.Args {
		log.Println(index, ":", arg)
	}
}
