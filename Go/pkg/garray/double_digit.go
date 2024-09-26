package garray

import (
	"fmt"
	"sort"
)

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
