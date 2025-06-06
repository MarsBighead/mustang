package gstack

import (
	"fmt"
	"sort"
	"strings"
)

// 14. 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix := lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := len(str1)
	if len(str2) < length {
		length = len(str2)
	}
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {

	n := len(s)
	if n%2 == 1 {
		return false
	}
	var (
		stack = []byte(s)
		top   = stack[n-1]
		tmp   = []byte{top}
	)

	var valid bool
	//fmt.Println(string(stack))
	stack = stack[:n-1]

	for i := n - 2; i >= 0; i-- {
		current := stack[len(stack)-1]
		if len(tmp) > 0 {
			top = tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) > 1 {
				top = current
				stack = stack[:len(stack)-1]
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				i--
			} else {
				break
			}
		}
		fmt.Println("stack", string(stack), "tmp:", string(tmp))
		p := string([]byte{current, top})
		fmt.Printf("n=%d, %v, %v, pair: %s\n", i, string(stack), string(tmp), p)
		if !(p == "()" || p == "{}" || p == "[]") {
			tmp = append(tmp, top, current)
		}

	}
	fmt.Println("End:", string(stack), "tmp:", string(tmp))
	if len(tmp) == 0 && len(stack) == 0 {
		valid = true
	}
	return valid

}

// 32. 最长有效括号
// https://leetcode.cn/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	cnt, n := 0, len(s)
	if n < 2 {
		return cnt
	}
	stack := []int{-1}
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				cnt = max(cnt, i-stack[len(stack)-1])
			}
		}

	}
	return cnt
}

func longestValidParenthesesv1(s string) int {
	cnt, n := 0, len(s)
	if n < 2 {
		return cnt
	}
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	left, right := 0, 0
	for _, ch := range s {
		if ch == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			cnt = max(cnt, 2*left)
		}
		if right > left {
			right, left = 0, 0
		}

	}
	fmt.Println("left", cnt)
	left, right = 0, 0
	for i := n - 1; i > 0; i-- {
		ch := s[i]
		if ch == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			cnt = max(cnt, 2*left)
		}
		if right < left {
			right, left = 0, 0
		}

	}
	fmt.Println("right", cnt)
	return cnt
}

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string]int)
	ans := [][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		if idx, ok := mp[sortedStr]; ok {
			ans[idx] = append(ans[idx], str)
		} else {
			ans = append(ans, []string{str})
			mp[sortedStr] = len(ans) - 1
		}
	}

	return ans
}

// 71. 简化路径
// https://leetcode.cn/problems/simplify-path/description/
func simplifyPath(path string) string {
	stack := []string{}
	for _, dir := range strings.Split(path, "/") {
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		} else if dir != "" && dir != "." {
			stack = append(stack, dir)
		}
	}
	return "/" + strings.Join(stack, "/")
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
