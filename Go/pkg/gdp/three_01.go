package gdp

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/description/?envType=study-plan-v2&envId=dynamic-programming
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	r0, r1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		r0, r1 = r1, max(r0+nums[i], r1)
	}
	return r1
}
