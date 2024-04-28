package bsearch

import "fmt"

// search https://leetcode.cn/problems/binary-search/
func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	if right == 0 && nums[right] == target {
		return right
	}
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("mid index=", mid, right)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
