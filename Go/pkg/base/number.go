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

// 89. 格雷编码
// https://leetcode.cn/problems/gray-code/description/
func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
		fmt.Printf("%.4b^%.4b=%.4b,%d\n", i>>1, i, ans[i], i)
	}

	return ans

}

// 191. 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/description/
func hammingWeight(n int) int {
	ones := 0
	for ; n > 0; n &= n - 1 {
		ones++
	}
	return ones
}

// 371. 两整数之和
// 给你两个整数 a 和 b ，不使用 运算符 + 和 - ​​​​​​​，计算并返回两整数之和。-1000 <= a, b <= 1000
// https://leetcode.cn/problems/sum-of-two-integers/description/
func getSum(a int, b int) int {

	for b != 0 {
		carry := a & b << 1
		a ^= b
		b = carry
	}
	return a
}

// 190. 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/description/
func reverseBits(n uint32) (rev uint32) {
	for i := 0; i < 32 && n > 0; i++ {
		rev |= n & 1 << (31 - i)
		n >>= 1
	}
	return
}
