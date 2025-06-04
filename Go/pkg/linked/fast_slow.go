package linked

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
