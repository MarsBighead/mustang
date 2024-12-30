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
	for i := 1; i < idx; i++ {
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

func reverseWordsv1(s string) string {
	n := len(s)
	if s[0] != ' ' && n == 1 {
		return s
	}
	step, idx := 0, n
	ss := make([]byte, n)
	for i := 0; i < n; i++ {
		if s[i] != ' ' {
			step++
		} else {
			if step > 0 {
				idx = idx - step
				copy(ss[idx:], s[i-step:i])
				idx--
				ss[idx] = ' '
			}
			step = 0
		}
		if i == n-1 && step > 0 {
			//fmt.Printf("%d: idx=%d,step=%d, %s\n", n, idx, step, s[n-step:])
			// case: "EPY2giL"
			if idx < n {
				ss[idx] = ' '
			}
			idx = idx - step
			copy(ss[idx:], s[n-step:])
			idx--
		}
	}
	return string(ss[idx+1:])
}

// 557. 反转字符串中的单词 III
// https://leetcode.cn/problems/reverse-words-in-a-string-iii/description/
func reverseWords3(s string) string {
	if len(s) <= 1 {
		return s
	}
	step, length := 0, len(s)
	ss := []byte(s)
	for i := 0; i < length; i++ {
		if s[i] != ' ' {
			step++
		} else {
			if step > 1 {
				for left, right := i-step, i-1; left < right; left++ {
					ss[left], ss[right] = ss[right], ss[left]
					right--
				}
			}
			step = 0
		}
	}
	if step > 1 {
		for left, right := length-step, length-1; left < right; left++ {
			ss[left], ss[right] = ss[right], ss[left]
			right--
		}
	}
	return string(ss)
}
