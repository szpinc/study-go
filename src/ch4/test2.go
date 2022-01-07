package main

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func rotate(s []int, target int) []int {
	length := len(s)
	s2 := make([]int, length)
	for i := 0; i < length; i++ {
		index := i + target
		if index > length {
			index = index - length
		}
		s2[i] = s[index]
	}
	return s2
}
