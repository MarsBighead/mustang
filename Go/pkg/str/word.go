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

// 6. Z 字形变换
// https://leetcode.cn/problems/zigzag-conversion/description
/*
方法三：直接构造
我们来研究方法一中矩阵的每个非空字符会对应到 s 的哪个下标（记作 idx），从而直接构造出答案。
由于 Z 字形变换的周期为 t=2r−2，因此对于矩阵第一行的非空字符，其对应的 idx 均为 t 的倍数，即 idx≡0(modt)；同理，对于矩阵最后一行的非空字符，应满足 idx≡r−1(modt)。
对于矩阵的其余行（行号设为 i），每个周期内有两个字符，第一个字符满足 idx≡i(modt)，第二个字符满足 idx≡t−i(modt)
*/
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	t := numRows*2 - 2
	result := make([]byte, 0, n)
	for i := 0; i < numRows; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			result = append(result, s[j+i]) // 当前周期的第一个字符
			if i > 0 && i < numRows-1 && j+t-i < n {
				result = append(result, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(result)

}
