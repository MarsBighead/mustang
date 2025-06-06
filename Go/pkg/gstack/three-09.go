package gstack

import (
	"fmt"
	"sort"
)

// 973. 最接近原点的 K 个点
// https://leetcode.cn/problems/k-closest-points-to-origin/description/
func kClosest(points [][]int, k int) [][]int {
	if len(points) == 1 && k == 1 {
		return points
	}
	distances := []int{}
	hash := make(map[int][]int)
	for idx, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		if len(hash[distance]) == 0 {
			hash[distance] = []int{idx}
			fmt.Println(hash[distance])
			distances = append(distances, distance)
		} else {
			hash[distance] = append(hash[distance], idx)
		}
	}
	fmt.Println(distances, hash, k)
	sort.Ints(distances)
	ans := [][]int{}
	for i := 0; i < k; {
		for _, idx := range hash[distances[i]] {
			if len(ans) >= k {
				break
			}
			ans = append(ans, points[idx])
			i++
		}

	}
	return ans[:k]
}

// 946. 验证栈序列
// https://leetcode.cn/problems/validate-stack-sequences/description/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	ok := true
	n, m := 0, 0
	for _, pu := range pushed {
		stack = append(stack, pu)
		n = len(stack)
		if stack[n-1] == popped[m] {
			stack = stack[:n-1]
			m++
			n--
		}
		for n > 0 {
			if stack[n-1] == popped[m] {
				stack = stack[:n-1]
				m++
				n--
			} else {
				break
			}
		}
	}
	fmt.Println("pushed stack", stack, popped)
	n = len(stack)
	for n > 0 {
		if stack[n-1] == popped[m] {
			m++
			n--
		} else {
			ok = false
			break
		}
	}
	fmt.Println("popped stack", stack, popped)

	return ok
}

// 973. 最接近原点的 K 个点
// https://leetcode.cn/problems/k-closest-points-to-origin/description/
func kClosestv1(points [][]int, k int) [][]int {
	if len(points) == 1 && k == 1 {
		return points
	}
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}
