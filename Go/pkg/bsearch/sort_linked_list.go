package bsearch

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
			current.Next = &ListNode{Val: second.Val}
			second = second.Next
		} else {
			current.Next = &ListNode{Val: first.Val}
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		current            = head
		prev, result, back *ListNode
	)
	prev = &ListNode{Val: head.Val}
	result = prev
	for current != nil {
		back = current
		prev = current
		for current != nil {
			if prev.Val > current.Val {
				prev = &ListNode{Val: current.Val}
				back = back.Next
				current = back.Next
			}
			current = current.Next
		}
		prev = prev.Next
		current = back.Next
	}

	return result
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
