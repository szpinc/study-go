package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	m := make(map[string]bool)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10 && scanner.Scan(); i++ {

		text := scanner.Text()

		if !m[text] {
			m[text] = true
		}
	}

	fmt.Println("--------------------------")

	for k := range m {
		fmt.Println(k)
	}
}
