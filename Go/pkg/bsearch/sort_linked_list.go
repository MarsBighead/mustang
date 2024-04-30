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
		head    = new(ListNode)
		current = head
	)
	head.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{
			Val: nums[i],
		}
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
