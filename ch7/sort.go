package main

import (
	"fmt"
	"sort"
)

func main() {

	list := StringList{1, 34, 56, 123, 2, 4}

	fmt.Println("list排序前:", list)

	sort.Sort(&list)

	fmt.Println("list排序后:", list)
}

type StringList []int

func (str *StringList) Len() int {
	return len(*str)
}

func (str *StringList) Less(i, j int) bool {
	return (*str)[i] < (*str)[j]
}

func (str *StringList) Swap(i, j int) {
	(*str)[i], (*str)[j] = (*str)[j], (*str)[i]
}
