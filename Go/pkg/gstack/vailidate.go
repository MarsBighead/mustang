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

// 85. 最大矩形
// https://leetcode.cn/problems/maximal-rectangle/description/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		if matrix[0][0] == '1' {
			return 1
		} else {
			return 0
		}
	}
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	ans := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	aa := make([][]int, m)
	for i, row := range matrix {
		dp[i] = make([]int, n)
		aa[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}

	for x := 0; x < n; x++ {
		up := make([]int, m)
		down := make([]int, m)
		stack := []int{}
		// cell值为0，则上界为-1
		// 值不为0，X方向长度相等时候，下界上界为最早的一个x长度相等的cell Y-1
		// ?? 可能描述不准确
		for y, row := range dp {
			n := len(stack)
			for n > 0 && dp[stack[n-1]][x] >= row[x] {
				stack = stack[:n-1]
				n--
			}
			up[y] = -1
			if n > 0 {
				up[y] = stack[n-1]
			}
			stack = append(stack, y)
		}
		//fmt.Println(x, "up  ", up)
		stack = []int{}
		// cell值为0，则下界为m
		// cell值为1，X方向长度相等时候，下界上界为最早的一个x长度相等的cell Y
		for y := m - 1; y >= 0; y-- {
			n := len(stack)

			for n > 0 && dp[stack[n-1]][x] >= dp[y][x] {
				stack = stack[:n-1]
				n--
			}
			down[y] = m
			if n > 0 {
				down[y] = stack[n-1]
			}
			stack = append(stack, y)
		}
		//fmt.Println(x, "down", down)
		for y, row := range dp {
			height := down[y] - up[y] - 1
			area := height * row[x]
			aa[y][x] = area
			ans = max(ans, area)
		}

		//fmt.Println()

	}

	for _, row := range dp {
		fmt.Println(row)
	}
	fmt.Println()
	for _, arow := range aa {
		fmt.Println(arow)
	}
	return ans
}
