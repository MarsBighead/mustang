package garray

import "container/heap"

type Pair struct{ i, j int }

type Hp struct {
	data         []Pair
	nums1, nums2 []int
}

// sort.Interface
func (h Hp) Len() int { return len(h.data) }
func (h Hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h Hp) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

// heap.Interface
/*
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
*/
func (h *Hp) Push(v interface{}) { h.data = append(h.data, v.(Pair)) }
func (h *Hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/?envType=study-plan-v2&envId=top-interview-150
// 373. 查找和最小的 K 对数字
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	h := Hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, Pair{i, 0})
	}
	var ans [][]int
	for h.Len() > 0 && len(ans) < k {
		// heap.Pop and heap.Push ensure data keep min-heap sort
		p := heap.Pop(&h).(Pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, Pair{i, j + 1})
		}
	}
	return ans
}
