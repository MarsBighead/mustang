package gstack

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
// 删除字符串中相邻重复的字符
func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
