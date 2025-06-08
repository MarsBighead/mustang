package gdp

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=dynamic-programming
//746. 使用最小花费爬楼梯

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	s1, s2 := 0, 0
	for i := 2; i <= n; i++ {
		fmt.Println("s2=", s2, s2+cost[i-1], s1+cost[i-2])
		s1, s2 = s2, min(s2+cost[i-1], s1+cost[i-2])
	}
	return s2

}

// https://leetcode.cn/problems/delete-and-earn/description/?envType=study-plan-v2&envId=dynamic-programming
// 740. 删除并获得点数
func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	ans := 0
	sum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val == nums[i-1] {
			sum[len(sum)-1] += val
		} else if val == nums[i+1] {
			sum = append(sum, val)
		} else {
			ans += rob(sum)
			sum = []int{val}
		}
	}
	ans += rob(sum)
	return ans
}
