package number

import "sort"

// https://leetcode.cn/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-interview-150
// 56. 合并区间
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{
		intervals[0],
	}
	for i := 1; i < n; i++ {
		interval := ans[len(ans)-1]
		if intervals[i][0] <= interval[1] {
			if intervals[i][1] > interval[1] {
				interval[1] = intervals[i][1]
			}

		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
