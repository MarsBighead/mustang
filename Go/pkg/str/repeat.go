package str

import "fmt"

// 459. 重复的子字符串
// https://leetcode.cn/problems/repeated-substring-pattern/description/
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}
	ds := s + s
	ok := false
	for i := 1; i < n; i++ {
		if ds[i:i+n] == s {
			ok = true
			fmt.Printf("%d, n=%d, sub string: %s\n", i, n, ds[i:i+n])
			break
		}
	}

	return ok
}
