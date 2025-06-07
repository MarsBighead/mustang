package linked

import "sort"

// LCR 123. 图书整理 I
// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/description/
func reverseBookList(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	current := head
	cnt := 0
	for current.Next != nil {
		cnt++
		current = current.Next
	}
	ans := make([]int, cnt+1)
	current = head
	for current.Next != nil {
		ans[cnt] = current.Val
		current = current.Next
		cnt--
	}
	ans[cnt] = current.Val

	return ans
}
func reverseBookListv1(head *ListNode) []int {

	if head == nil {
		return []int{}
	}
	var (
		prev *ListNode
		cnt  int
	)

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
		cnt++
	}

	ans := make([]int, cnt)
	cnt = 0
	for head != nil {
		ans[cnt] = head.Val
		head = head.Next
		cnt++
	}
	return ans
}

// https://leetcode.cn/problems/majority-element/description/
// 169. 多数元素
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore 算法
// Reference: https://leetcode.cn/problems/majority-element/solutions/1/169-duo-shu-yuan-su-mo-er-tou-piao-qing-ledrh/
func majorityElementv1(nums []int) int {
	candidate, count := -1, 0
	for _, num := range nums {
		if num == candidate {
			count++
		} else {
			if count == 0 {
				candidate = num
				count = 1
			} else {
				count--
			}
		}
	}
	return candidate

}
