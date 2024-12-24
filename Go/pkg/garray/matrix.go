package garray

// 54.螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix[0]) == 0 {
		return []int{}
	}
	x, y := len(matrix[0]), len(matrix)
	r := make([]int, x*y)
	index := 0
	left, right, top, bottom := 0, x-1, 0, y-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			r[index] = matrix[top][i]
			index++
		}
		for j := top + 1; j <= bottom; j++ {
			r[index] = matrix[j][right]
			index++
		}
		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				r[index] = matrix[bottom][i]
				index++
			}
			for j := bottom; j > top; j-- {
				r[index] = matrix[j][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--

	}
	return r
}
