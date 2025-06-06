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

// 57. 插入区间
// https://leetcode.cn/problems/insert-interval/?envType=study-plan-v2&envId=top-interview-150
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	start, end := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > end {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{start, end})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < start {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			if interval[0] < start {
				start = interval[0]
			}
			if interval[1] > end {
				end = interval[1]
			}
		}

	}

	if !merged {
		ans = append(ans, []int{start, end})
	}
	return ans

}

// https://leetcode.cn/problems/sqrtx/description
// 69. x 的平方根
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
