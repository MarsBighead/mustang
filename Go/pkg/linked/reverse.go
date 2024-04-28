package linked

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	return reverse(cur, head)
}

func reverse(prev, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	} else {
		node := cur.Next
		cur.Next = prev
		return reverse(cur, node)
	}

}

func ArrayToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head

}

func PrintList(head *ListNode) {
	if head != nil {
		fmt.Print(head.Val, " ")
	}
	for head.Next != nil {
		head = head.Next
		fmt.Print(head.Val, " ")
	}
	fmt.Println()
}
