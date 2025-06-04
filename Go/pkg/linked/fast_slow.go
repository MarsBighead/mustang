package linked

import "sort"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
// 删除排序数组中的重复项 II
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	left := 2
	for i := 2; i < n; i++ {
		if nums[left-2] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

// https://leetcode.cn/problems/majority-element/description/
// 169. 多数元素
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore 算法
// Reference: https://leetcode.cn/problems/majority-element/solutions/1/169-duo-shu-yuan-su-mo-er-tou-piao-qing-ledrh/
func majorityElementv1(nums []int) int {
	candidate, count := -1, 0
	for _, num := range nums {
		if num == candidate {
			count++
		} else {
			if count == 0 {
				candidate = num
				count = 1
			} else {
				count--
			}
		}
	}
	return candidate

}
