package garray

// 724. 寻找数组的中心下标
// https://leetcode.cn/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	sum := 0
	for i, v := range nums {
		sum += v
		nums[i] = sum
	}
	for i, v := range nums {
		if i > 0 {
			if sum-v == nums[i-1] {
				return i
			}
		} else {
			if sum-v == 0 {
				return 0
			}
		}
	}
	return -1
}
