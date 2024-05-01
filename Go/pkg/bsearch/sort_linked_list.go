package bsearch

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	}
	n := len(lists)
	return mergeTwoLists(mergeKLists(lists[0:n/2]), mergeKLists(lists[n/2:]))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var (
		head          = &ListNode{}
		first, second = list1, list2
	)

	current := head
	//current = &ListNode{Val: list1.Val}
	for first != nil && second != nil {
		if first.Val > second.Val {
			current.Next = second
			second = second.Next
		} else {
			current.Next = first
			first = first.Next
		}

		current = current.Next
	}
	if first != nil {
		current.Next = first
	}
	if second != nil {
		current.Next = second
	}
	return head.Next
}

func array2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var (
		//head    = new(ListNode)
		head    = &ListNode{Val: nums[0]}
		current = head
	)
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	//current = nil
	return head
}

func AddCycle2List(head *ListNode, n int) *ListNode {
	if n == -1 || head == nil {
		return head
	}

	var (
		i       int
		current = head
		tail    *ListNode
	)
	for current.Next != nil {
		current = current.Next
	}
	tail = current
	current = head

	for current != nil {
		if i == n {
			tail.Next = current
			break
		}
		i += 1
		current = current.Next
	}
	return head
}

func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println()
		return
	}
	var (
		current = head
	)
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// https://leetcode.cn/problems/sort-list/description/
// 归并排序
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		slow, fast = head, head
		prev       *ListNode
	)
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := head, slow
	prev.Next = nil
	PrintList(left)
	PrintList(right)

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	for left != nil && right != nil {
		if left.Val <= right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}
	if left != nil {
		current.Next = left
	}
	if right != nil {
		current.Next = right
	}
	return result.Next

}

// https://leetcode.cn/problems/sort-list/description/
// 冒泡法
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		out  = head
		tail *ListNode
	)
	for out != nil {
		in := head
		for in != nil && in.Next != tail {
			if in.Val > in.Next.Val {
				in.Val, in.Next.Val = in.Next.Val, in.Val
			}
			in = in.Next
		}
		tail = in
		out = out.Next
	}
	return head
}

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}

	slow, fast := head, head
	PrintList(fast)
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	PrintList(slow)
	PrintList(fast)
	return false

}

func mergeKListsV2(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	}
	var (
		head   = lists[0]
		length = len(lists)
		tail   *ListNode
	)

	for i := 0; i < length-1; i++ {
		if head == nil {
			head = lists[i]
		}
		current := lists[i]
		for current != nil {
			if current.Next == nil {
				tail = current
			}
			current = current.Next
		}
		if tail != nil {
			tail.Next = lists[i+1]
		}
	}
	PrintList(head)
	mergeSort(head)
	return head
}
