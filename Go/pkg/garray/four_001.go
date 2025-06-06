package garray

import "fmt"

// 1095. 山脉数组中查找目标值
// https://leetcode.cn/problems/find-in-mountain-array/description/
func findInMountainArray(target int, mountainArr []int) int {
	left, right, length := 0, len(mountainArr)-1, len(mountainArr)

	for left < right {
		mid := left + (right-left)/2
		v, t := mountainArr[mid], mountainArr[mid+1]
		if v > t {
			right = mid
		} else {
			left = mid + 1
		}
	}
	idx := findTarget(0, left, target, mountainArr)
	fmt.Println("peak=", left)
	if idx == -1 {
		idx = findTarget(left+1, length-1, target, mountainArr)
	}
	return idx
}

// findTarget need to take sort desc or asc in situation
func findTarget(left, right, target int, mountainArr []int) int {
	isRight := true
	if left == 0 {
		isRight = false
	}
	idx := -1
	fmt.Println(mountainArr[left:right+1], "xxx", left, right)
	for left <= right {
		mid := (right + left) / 2
		v := mountainArr[mid]
		fmt.Println("left:right", left, right, mid, v, target)
		if v == target {
			return mid
		} else if v > target {
			if isRight {
				left = mid + 1
			} else {
				right = mid - 1
			}

		} else {
			if isRight {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}
	}

	return idx

}
