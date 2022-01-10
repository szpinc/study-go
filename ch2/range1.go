package main

import "fmt"

func main() {

	s := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	maxIndex := len(s) - 1

	for i, v := range s {
		if i == maxIndex {
			s[0] += v
		} else {
			s[i+1] += v
		}
	}

	switch 2 + 4 {

	}

	fmt.Println(s)
}
