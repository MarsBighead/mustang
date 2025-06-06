package garray

import "regexp"

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
