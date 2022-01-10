package main

import "fmt"

func main() {
	s := []string{"1", "2", "3", "4", "4", "5"}
	fmt.Println("s:", s)
	s2 := removeRepeatString(s)
	fmt.Println("s2:", s2)
}

func removeRepeatString(s []string) []string {
	if len(s) == 0 {
		return s
	}

	length := len(s)
	for i, v := range s {
		if i == length-1 {
			continue
		}

	}

	return s
}

func remove(s []string) []string {
	i := len(s)
	for i, v := range s {
		
	}
}
