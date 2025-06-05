package str

import "strings"

// https://leetcode.cn/problems/length-of-last-word
// 58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	n := len(s)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			cnt++
		} else {
			return cnt
		}
	}
	return n
}
func lengthOfLastWordv1(s string) int {
	index := len(s) - 1
	ans := 0
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		index--
		ans++
	}
	return ans
}
