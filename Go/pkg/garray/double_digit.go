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
