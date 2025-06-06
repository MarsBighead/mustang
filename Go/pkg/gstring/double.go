package gstring

import (
	"regexp"
	"strings"
)

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
// 28. 找出字符串中第一个匹配项的下标
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

// https://leetcode.cn/problems/valid-anagram/description/?envType=study-plan-v2&envId=top-interview-150
// 42. 有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//s 和 t 仅包含小写字母
	hash := [26]int{}
	for _, a := range s {
		hash[a-'a']++
	}
	for _, a := range t {
		if hash[a-'a'] == 0 {
			return false
		}
		hash[a-'a']--
	}

	return true
}

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
