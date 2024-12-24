package garray

import "regexp"

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
