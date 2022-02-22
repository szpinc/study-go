package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = true
	}
	return nums[0]
}
