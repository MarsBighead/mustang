package gdp

import "fmt"

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
