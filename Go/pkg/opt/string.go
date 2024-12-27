package opt

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string]int)
	ans := [][]string{}
	for _, str := range strs {
		ss := []rune(str)
		sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
		fmt.Println(string(ss))
		if idx, ok := mp[string(ss)]; ok {
			ans[idx] = append(ans[idx], str)
		} else {
			ans = append(ans, []string{str})
			mp[string(ss)] = len(ans) - 1
		}
	}

	return ans
}

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string/description/
func reverseWords(s string) string {
	if s[0] != ' ' && len(s) == 1 {
		return s
	}
	s = " " + s + " "
	m, n, idx := 1, 0, len(s)
	ss := make([]byte, idx)
	for i := 1; i < len(ss); i++ {
		if s[i-1] == ' ' && s[i] != ' ' {
			m = i - 1
		}
		if s[i] != ' ' && s[i+1] == ' ' {
			n = i + 1
			idx = idx - (n - m)
			copy(ss[idx:], s[m:n])
		}
	}
	return string(ss[idx+1:])
}
