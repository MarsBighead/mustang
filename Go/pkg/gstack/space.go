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

func frequencySort(s string) string {
	if len(s) <= 2 {
		return s
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	hash := make(map[rune]int)
	m := 0
	for _, ch := range s {
		hash[ch]++
		m = max(m, hash[ch])
	}
	buckets := make([][]rune, m+1)
	for ch, n := range hash {
		buckets[n] = append(buckets[n], ch)
	}
	ans := make([]rune, 0, len(s))
	for i := m; i > 0; i-- {
		if len(buckets[i]) > 0 {
			for _, ch := range buckets[i] {
				for j := 0; j < i; j++ {
					ans = append(ans, ch)
				}

			}
		}
	}
	return string(ans)
}

type entity struct {
	ch    rune
	Count int
}

func frequencySortv1(s string) string {
	if len(s) <= 2 {
		return s
	}
	dict := make([]*entity, 62)
	for _, ch := range s {
		if ch-'A' >= 0 && ch-'Z' <= 0 {
			idx := 10 + ch - 'A'
			if dict[idx] == nil {
				dict[idx] = &entity{ch, 1}
			} else {
				dict[idx].Count++
			}

		} else if ch-'a' >= 0 && ch-'z' <= 0 {
			idx := 36 + ch - 'a'
			if dict[idx] != nil {
				dict[idx].Count++
			} else {
				dict[idx] = &entity{ch, 1}
			}

		} else if ch-'0' >= 0 && ch-'9' <= 0 {
			idx := ch - '0'
			if dict[idx] == nil {
				dict[idx] = &entity{ch, 1}
			} else {
				dict[idx].Count++
			}

		}
	}
	sort.Slice(dict, func(i, j int) bool {
		n, m := 0, 0
		if dict[i] != nil {
			m = dict[i].Count
		}
		if dict[j] != nil {
			n = dict[j].Count
		}
		return m > n
	})
	ans := []rune{}
	for _, d := range dict {
		if d != nil {
			fmt.Println(d.Count, d.ch)
			ss := []rune{}
			for i := 0; i < d.Count; i++ {
				ss = append(ss, d.ch)
			}
			ans = append(ans, ss...)
		}
	}
	return string(ans)
}
