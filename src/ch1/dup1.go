package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < 5 && input.Scan(); i++ {
		counts[input.Text()]++
	}

	for key, value := range counts {
		fmt.Printf("%s:%d\n", key, value)
	}
}
