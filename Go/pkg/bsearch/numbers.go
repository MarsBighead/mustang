package bsearch

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	MAX = 1<<31 - 1
	MIN = -2 << 30
)

// 8. 字符串转换整数 (atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/description/
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`^(?:[+-]{0,1})(\d+)`)
	nums := re.FindStringSubmatch(s)
	if len(nums) == 0 {
		return 0
	}
	//n, _ := strconv.Atoi(nums[0])
	var n, t int
	nss := []byte(nums[0])
	for i, ch := range nss {
		if nss[0] == '-' && i > 0 {
			t = n*10 - int((ch - 48))
			if t < MIN {
				return MIN
			}
			n = t
		} else if nss[0] == '+' && i > 0 {
			t = n*10 + int((ch - 48))
			if t > MAX {
				return MAX
			}
			n = t
		} else {
			if nss[0] != '-' && nss[0] != '+' {
				t = n*10 + int((ch - 48))
				if t > MAX {
					return MAX
				}
				n = t
			}
		}
	}
	//fmt.Println("n=", n)
	return n
}

// 91. 解码方法
// https://leetcode.cn/problems/decode-ways/description/
// s[i-2],s[i-1]:
// 1~26: 10-1,20-1,26-1,>26-0
// 11~25: 2
func numDecodings(s string) int {
	n := len(s)
	if n > 0 && s[0] == '0' {
		return 0
	}
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
		fmt.Println(a, b, c, s)
	}
	return c
}
