package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := make([]string, 10)
	for i := range s {
		s[i] = strconv.Itoa(i + 1)
	}
	fmt.Println("s:", s)
	s1 := s[:3]
	fmt.Println("s1:", s1)
	s2 := appendString(s1, "11", "12", "11", "12", "11", "12", "11", "12")
	fmt.Println("s2:", s2)
	fmt.Println("s:", s)
}

func appendString(s []string, str ...string) []string {
	// 切片总大小
	capcity := cap(s)
	// 切片长度
	length := len(s)
	// 原切片剩余空间
	remainLength := capcity - length

	var newSlice []string

	// 剩余空间充足，新增内容追加到原切片数组
	if remainLength >= len(str) {
		newSlice = s[:length+len(str)]
	} else {
		// 剩余空间不足，需要申请一个新的切片
		newSlice = make([]string, length+len(str))
		for i, v := range s {
			newSlice[i] = v
		}
	}

	for i, v := range str {
		newSlice[length+i] = v
	}
	return newSlice
}
