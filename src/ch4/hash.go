package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	hashType string
	help     bool
)

func init() {
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&hashType, "t", "hash256", "`加密类型`:hash256,hash384,hash512")
}

func main() {

	value := flag.Arg(len(flag.Args()) - 1)

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	println("type:", hashType)
	println("value:", value)

	switch hashType {
	case "hash256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(value)))
		break
	case "hash384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(value)))
		break
	case "hash512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(value)))
		break
	default:
		fmt.Errorf("hash type invalid")
	}
}
