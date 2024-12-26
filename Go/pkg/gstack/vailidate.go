package gstack

import "fmt"

// 946. 验证栈序列
// https://leetcode.cn/problems/validate-stack-sequences/description/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	ok := true
	n, m := 0, 0
	for _, pu := range pushed {
		stack = append(stack, pu)
		n = len(stack)
		if stack[n-1] == popped[m] {
			stack = stack[:n-1]
			m++
			n--
		}
		for n > 0 {
			if stack[n-1] == popped[m] {
				stack = stack[:n-1]
				m++
				n--
			} else {
				break
			}
		}
	}
	fmt.Println("pushed stack", stack, popped)
	n = len(stack)
	for n > 0 {
		if stack[n-1] == popped[m] {
			m++
			n--
		} else {
			ok = false
			break
		}
	}
	fmt.Println("popped stack", stack, popped)

	return ok
}

// 3. 单调栈模板
// 3.1 单调递增栈模板
func monotoneIncreasingStack(nums []int) []int {
	stack := []int{}
	for _, num := range nums {
		n := len(stack)
		for n > 0 && num >= stack[n-1] {
			stack = stack[:n-1]
		}
		stack = append(stack, num)
	}
	return stack
}

// 3.2 单调递减栈模板
func monotoneDecreasingStack(nums []int) []int {
	stack := []int{}
	for _, num := range nums {
		n := len(stack)
		for n > 0 && num <= stack[n-1] {
			stack = stack[:n-1]
		}
		stack = append(stack, num)
	}
	return stack
}
