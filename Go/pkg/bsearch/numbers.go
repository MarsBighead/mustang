package bsearch

import (
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
