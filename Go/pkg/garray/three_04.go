package garray

import (
	"fmt"
	"regexp"
)

// 468.验证IP地址
// https://leetcode.cn/problems/validate-ip-address/description/
func ValidIPAddress(queryIP string) string {
	var (
		ipv4 = regexp.MustCompile("^(?:(?:0\\.){3}0)|^((?:25[0-5]|2[0-4]\\d|((1\\d{2})|[1-9]?\\d|))((?:\\.)(?:25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))){3})$")
		ipv6 = regexp.MustCompile("^([0-9a-fA-F]{1,4})(\\:[0-9a-fA-F]{1,4}){7}$")
	)
	if ipv4.MatchString(queryIP) {
		return "IPv4"
	} else if ipv6.MatchString(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

// 498. 对角线遍历
// https://leetcode.cn/problems/diagonal-traverse/description/
func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 1 {
		return mat[0]
	}
	m, n := len(mat), len(mat[0])
	ans := []int{}
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			x := max(i-n+1, 0)
			y := min(i, n-1)
			fmt.Println("1:", x, y)
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(i-m+1, 0)
			fmt.Println("0:", x, y)
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}
