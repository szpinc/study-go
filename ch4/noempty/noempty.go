package main

import "fmt"

func main() {
	s := []string{"1", "", "3", "", "4"}
	fmt.Println("s:", s)
	s2 := removeEmpty(s)
	fmt.Println("s2:", s2)
}

func removeEmpty(s []string) []string {
	index := 0
	for _, str := range s {
		if str != "" {
			s[index] = str
			index++
		}
	}
	return s[:index]
}
