package garray

import (
	"math"
)

// 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/?envType=study-plan-v2&envId=top-interview-150
// 滑动窗口方法(sliding window)，是对暴力法的优化，保证运算不超时
func minSubArrayLen(target int, nums []int) int {
	ans, n := math.MaxInt, len(nums)
	start, end, x := 0, 0, 0
	for ; end < n; end++ {
		x += nums[end]
		for x >= target {
			if ans > end-start+1 {
				ans = end - start + 1
			}
			x -= nums[start]
			start++
		}

	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

// https://leetcode.cn/problems/h-index/description
// 274. H 指数
func hIndex(citations []int) int {
	n := len(citations)
	hs := make([]int, n+1)
	for _, citation := range citations {
		if citation > n {
			citation = n
		}
		hs[citation] += 1
	}
	for i, total := n, 0; i >= 0; i-- {
		total += hs[i]
		if total >= i {
			return i
		}
	}

	return 0

}

// https://leetcode.cn/problems/game-of-life/
// 289. 生命游戏
func GameOfLife(board [][]int) {
	//00:死， 01:活变死， 10死变活；11:活
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	x, y := len(board[0]), len(board)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			num := getLiveNum(i, j, board)
			if board[i][j] == 1 {
				if num == 2 || num == 3 {
					board[i][j] = board[i][j] | 2
				}
			} else {
				if num == 3 {
					board[i][j] = board[i][j] | 2
				}
			}
		}
	}
	for i := 0; i < y; i++ {
		//fmt.Println(board[i])
		for j := 0; j < x; j++ {
			board[i][j] >>= 1
		}
	}
}

func getLiveNum(i, j int, board [][]int) int {
	num := 0
	x, y := len(board[0]), len(board)
	postions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, pos := range postions {
		pos[0] += i
		pos[1] += j
		if pos[0] < 0 || pos[0] >= y || pos[1] < 0 || pos[1] >= x {
			continue
		}
		num += board[pos[0]][pos[1]] & 1
	}

	return num
}
