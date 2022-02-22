package main

func main() {

	testCase := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	maxArea := maxArea(testCase)

	println(maxArea)
}

func maxArea(height []int) int {
	// 定义左、右指针位置
	left, right := 0, len(height)-1
	// 定义最大面积
	maxArea, area := 0, 0

	for left < right {
		if height[left] > height[right] {
			area = (right - left) * height[right]
			right--
		} else {
			area = (right - left) * height[left]
			left++
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
