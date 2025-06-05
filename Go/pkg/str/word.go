package str

import (
	"regexp"
	"strings"
)

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

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	re := regexp.MustCompile(needle)
	indices := re.FindStringIndex(haystack)
	if indices != nil {
		return indices[0]
	}
	return -1
}

func strStrv1(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m > n {
		return -1
	}

	for i := 0; i < n; i++ {
		if haystack[i] == needle[0] {
			if i+m <= n {
				cap := haystack[i : i+m]
				if cap == needle {
					return i
				}
			} else {
				return -1
			}
		}
	}
	return -1
}
