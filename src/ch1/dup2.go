package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)

	files := os.Args

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		fmt.Printf("key:%s,value:%d", k, v)
	}

}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
