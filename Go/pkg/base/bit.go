package base

// https://leetcode.cn/problems/add-binary/description/?envType=study-plan-v2&envId=top-interview-150
// 67. 二进制求和
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	max := m
	if n > m {
		max = n
	}
	rn := max + 1
	ans := make([]byte, rn)
	carry := 0
	for i := 0; i < max; i++ {
		if i < m {
			carry += int(a[m-i-1] - '0')
		}
		if i < n {
			carry += int(b[n-i-1] - '0')
		}
		if carry%2 == 0 {
			ans[rn-i-1] = '0'
		} else {
			ans[rn-i-1] = '1'
		}
		carry /= 2
	}
	if carry > 0 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
