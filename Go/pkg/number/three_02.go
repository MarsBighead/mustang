package number

import "fmt"

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/?envType=study-plan-v2&envId=top-interview-150
// 201. 数字范围按位与
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		fmt.Printf("Pre: %d&%d,", right, right-1)
		right &= (right - 1)
		fmt.Printf("follow: %d\n", right)

	}
	return right
}
