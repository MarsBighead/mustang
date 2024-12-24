package garray

import (
	"regexp"
	"strconv"
	"strings"
)

func NextPermutation(nums []int) {

	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

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

// 165. 比较版本号
// https://leetcode.cn/problems/compare-version-numbers/description/
func compareVersion(version1 string, version2 string) int {
	var (
		vs1 = strings.Split(version1, ".")
		vs2 = strings.Split(version2, ".")
	)
	m, n := len(vs1), len(vs2)
	length := n
	if m > n {
		length = m
	}
	var (
		a, b, i int
	)
	for ; i < length; i++ {
		if i < m {
			a, _ = strconv.Atoi(vs1[i])
		} else {
			a = 0
		}
		if i < n {
			b, _ = strconv.Atoi(vs2[i])
		} else {
			b = 0
		}
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}
