package hash

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/description/
// 位图算法
func findDuplicates(nums []int) []int {
	length := len(nums)
	count := make([]int, length)
	for i := 0; i < length; i++ {
		v := nums[i]
		count[v-1] += 1
	}
	nums = nums[:0]
	for i, v := range count {
		if v == 2 {
			nums = append(nums, i+1)
		}
	}
	return nums
}
