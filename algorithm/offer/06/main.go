package main

import "sort"

func main() {
	testCase := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(testCase)
	println(result)
}

func threeSum(nums []int) [][]int {

	result := [][]int{}

	// 1.校验数组是否为空,长度是否小于3
	if nums == nil || len(nums) < 3 {
		return result
	}

	// 2.排序
	sort.Ints(nums)

	for i, num := range nums {
		// 大于0的数直接结束
		if num > 0 {
			return result
		}
		// 有相等的数
		if i > 0 && num == nums[i-1] {
			continue
		}
		// 定义左右指针
		left, right := i+1, len(nums)-1

		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			sum := num + nums[left] + nums[right]
			// sum > 0 需要将右指针左移
			if sum > 0 {
				right--
			} else if sum < 0 { // sum > 0 需要将左指针右移
				left++
			} else { // sum == 0 放进结果集中，同时移动左右指针
				result = append(result, []int{num, nums[left], nums[right]})
				right--
				left++
			}
		}
	}
	return result
}
