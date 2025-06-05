package garray

// https://leetcode.cn/problems/container-with-most-water/description/
// 11. 盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans := 0
	for right > left {
		area := min(height[left], height[right]) * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}
