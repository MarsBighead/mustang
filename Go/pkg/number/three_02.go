package number

import (
	"fmt"
	"math"
)

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

// 279. 完全平方数
// https://leetcode.cn/problems/perfect-squares/description/?envType=study-plan-v2&envId=top-100-liked
// 四平方和定理
func numSquares(n int) int {
	isSquare := func(x int) bool {
		y := int(math.Sqrt(float64(x)))
		return x == y*y
	}
	//4^k*(8*m+7)
	is4Squares := func(x int) bool {
		for x%4 == 0 {
			x /= 4
		}
		return x%8 == 7
	}
	if isSquare(n) {
		return 1
	}
	if is4Squares(n) {
		return 4
	}
	for i := 1; i*i < n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}
