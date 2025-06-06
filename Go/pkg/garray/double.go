package garray

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/container-with-most-water/description/
// 11. 盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans := 0
	for right > left {
		area := min(height[left], height[right]) * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

// https://leetcode.cn/problems/combination-sum/
// 39. Combination Sum(组合总和)
func CombinationSum(candidates []int, target int) (ans [][]int) {
	var (
		tmp []int
		n   int
	)
	for _, val := range candidates {
		if val < target {
			tmp = append(tmp, val)
		} else if val == target {
			ans = append(ans, []int{val})
		}
	}
	candidates = tmp
	tmp = []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {

		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			tmp = append(tmp, candidates[idx])
			dfs(target-candidates[idx], idx)
			tmp = tmp[:len(tmp)-1]
		} else {
			// 剪支无法避免递归数量增加
			return
		}
		n++
		fmt.Println(tmp)
	}
	dfs(target, 0)
	fmt.Println("n=", n)
	return
}

func CombinationSumV1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var (
		n   int
		row []int
		dfs func(target, idx int)
	)
	dfs = func(target, idx int) {
		n++
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), row...))
			return
		}
		for i := idx; i < len(candidates) && target > 0; i++ {
			// 剪枝逻辑
			if target-candidates[i] < 0 {
				break
			}
			row = append(row, candidates[i])

			dfs(target-candidates[i], i)
			row = row[:len(row)-1]
		}
		//fmt.Println(row)
	}

	dfs(target, 0)
	//fmt.Println("n=", n)
	return
}

// https://leetcode.cn/problems/combination-sum-ii/description/
// 40. Combination Sum(组合总和)

func CombinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	fmt.Println(candidates)
	var (
		row []int
		dfs func(target, idx int)
	)
	dfs = func(target, idx int) {
		if target == 0 {
			fmt.Println("    idx=", idx, row)
			ans = append(ans, append([]int(nil), row...))
			return
		}

		for i := idx; i < len(candidates) && target > 0; i++ {
			// 剪枝逻辑
			if target-candidates[i] < 0 {
				return
			}
			// skip idx=0, i=1, and it will not be put in recursive
			if i > idx && candidates[i] == candidates[i-1] {
				fmt.Println("skip", idx, i, row, candidates[i])
				continue
			}
			fmt.Println("pos", idx, i, row, candidates[i])
			row = append(row, candidates[i])

			dfs(target-candidates[i], i+1)
			row = row[:len(row)-1]
			if len(row) == 0 {
				fmt.Println("empty row", idx, i, row, candidates[i])
			}

			/*
				if i > 0 && candidates[i] == candidates[i-1] {
					continue
				}
			*/
		}

	}
	dfs(target, 0)
	return
}

// https://leetcode.cn/problems/first-missing-positive/description/
// 41. 缺失的第一个正数
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func FirstMissingPositiveV1(nums []int) int {
	n := len(nums)
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
			//continue
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	fmt.Println("--", nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://leetcode.cn/problems/unique-paths/description/
// 62. 不同路径
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Printf("\n\n###\n\n")
	for i := 0; i < m; i++ {
		fmt.Printf("%#v\n", dp[i])
	}
	return dp[m-1][n-1]
}

/*
思路： 排列组合
因为机器到底右下角，向下几步，向右几步都是固定的，

比如，m=3, n=2，我们只要向下 1 步，向右 2 步就一定能到达终点。
所以有

	C^m−1
	  m+n−2
*/
func UniquePathsV1(m int, n int) int {
	if m > n {
		n, m = m, n
	}
	cn := int64(n + m - 2)
	cm := int64(m - 1)
	ans := int64(1)
	for i := int64(0); i < cm; i++ {
		ans = ans * (cn - i) / (i + 1)
	}
	return int(ans)
}

// https://leetcode.cn/problems/unique-paths-ii/description/
// 63. 不同路径 II
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	idx := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		idx[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				idx[j] = 0
				fmt.Printf("%d ", idx[j])
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				idx[j] += idx[j-1]
			}
			fmt.Printf("%d ", idx[j])
		}
		fmt.Println()
	}

	fmt.Printf("%v\n", idx)
	return idx[n-1]
}

func UniquePathsWithObstaclesV1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {

			obstacleGrid[i][0] -= 1
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			obstacleGrid[0][j] -= 1
		}
	}
	// 需要修改支持低位的数字
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}
	fmt.Printf("###\n\n")
	for i := 0; i < m; i++ {
		fmt.Printf("%#v\n", obstacleGrid[i])
	}
	return 0 - obstacleGrid[m-1][n-1]
}
