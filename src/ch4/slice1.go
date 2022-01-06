package main

import "fmt"

// 切片是引用类型，底层使用数组
// 子切片共享底层数组
// 更改子切片会更改共享的底层数组
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(s)
	s2 := s[1:5]
	fmt.Println("s2:", s2)
	// 反转切片s2
	reverse(s2)
	// 此时s1的切面内容也会更改
	fmt.Println(s)

}

//反转切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
