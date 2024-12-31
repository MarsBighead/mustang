package base

import (
	"fmt"
	"strconv"
)

// 504. 七进制数
// https://leetcode.cn/problems/base-7/description/
func convertToBase7(num int) string {
	if num < 7 && num > -7 {
		return fmt.Sprintf("%d", num)
	}
	ans := ""
	isNagetive := false
	if num < 0 {
		num = -num
		isNagetive = true
	}
	for num > 0 {
		ans = strconv.Itoa(num%7) + ans
		num = num / 7
	}
	//fmt.Println(num)

	if isNagetive {
		ans = "-" + ans
	}
	return ans
}

func bitwiseComplement(n int) int {
	high := 0
	for i := 1; i <= 30; i++ {
		if n < 1<<i {
			break
		}
		high = i
	}
	mask := 1<<(high+1) - 1
	//fmt.Printf("mask: %b\n", mask)
	return n ^ mask
}

/*
0&0=0;
0&1=0;
1&0=0;
1&1=1;
*/
// 338. 比特位计数
// https://leetcode.cn/problems/counting-bits/description/
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
		fmt.Printf("%v\n", i&(i-1))
	}
	return bits

}
