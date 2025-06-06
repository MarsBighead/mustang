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

// 31. 下一个排列
// https://leetcode.cn/problems/next-permutation/description/
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

// https://leetcode.cn/problems/jump-game-ii/description
// 45. 跳跃游戏 II
// 正向查找可到达的最大位置，不计算最后一步
func jump(nums []int) int {
	var (
		pos   = 0
		steps = 0
		right = 0
	)
	for i := 0; i < len(nums)-1; i++ {
		if pos < i+nums[i] {
			pos = i + nums[i]
		}
		if i == right {
			right = pos
			steps++
		}
	}
	return steps
}

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

// https://leetcode.cn/problems/jump-game/description
// 55. 跳跃游戏
func canJump(nums []int) bool {
	n, right := len(nums), 0
	if nums[0] == 0 && n > 1 {
		return false
	}
	/*
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		for i, num := range nums {
			// i=0，计算right；若nums[0]=0,肯定false
			if i <= right {
				right = max(right, i+num)
				if right >= n-1 {
					return true
				}
			}
		}
		return false
	*/

	for i, num := range nums {
		if i <= right {
			fmt.Println("index:", i, right)
			if right < i+num {
				right = i + num
			}
			fmt.Println("changed index:", i, right)
		}
	}
	return right >= n-1
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

// https://leetcode.cn/problems/sort-colors/
// 75. 颜色分类
// 颜色排序，双指针
func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	for i, _ := range nums {
		if i < red {
			nums[i] = 0
		} else if i < red+white {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}

}
