package gstack

import "fmt"

// https://leetcode.cn/problems/daily-temperatures/description/
// 739. 每日温度
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
			fmt.Println("stack=", stack)
		}
		stack = append(stack, i)
		fmt.Println("ans=", i, stack)
	}
	return ans
}
