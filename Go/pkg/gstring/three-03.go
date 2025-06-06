package gstring

// https://leetcode.cn/problems/is-subsequence/description
// 392. 判断子序列
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	if m > n {
		return false
	}
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return m == i
}
