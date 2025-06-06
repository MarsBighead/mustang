package gstack

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/sort-characters-by-frequency/description/
// 451. 根据字符出现频率排序
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

// 496. 下一个更大元素 I
// https://leetcode.cn/problems/next-greater-element-i/description/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			hash[num] = stack[len(stack)-1]
		} else {
			hash[num] = -1
		}

		stack = append(stack, num)
	}
	for i, num := range nums1 {
		nums1[i] = hash[num]
	}
	return nums1
}

func nextGreaterElementv2(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		hash[nums2[i]] = i
	}
	for i, num := range nums1 {
		index := hash[num]
		found := false
		for j := index + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				nums1[i] = nums2[j]
				found = true
				break
			}
		}
		if !found {
			nums1[i] = -1
		}

	}
	return nums1
}
