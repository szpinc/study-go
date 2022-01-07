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

	var temp = s
	for true {
		t, stat := remove(temp)
		if stat {
			return t
		}
		temp = t
	}
	return s
}

func remove(s []string) ([]string, bool) {
	index := 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}
		if s[i] == s[i-1] {
			s[index] = s[i-1]
			i++
		}
	}
	if index == 0 {
		return s, true
	}
	return s[:index], false
}
