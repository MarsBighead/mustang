package bsearch

import "fmt"

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		result   = &ListNode{Next: head}
		unsorted = result.Next
		current  = head.Next
	)
	fmt.Println("Cal:")
	for current != nil {
		fmt.Println("Loop")
		PrintList(result.Next)
		if unsorted.Val <= current.Val {
			unsorted = unsorted.Next
		} else {
			prev := result
			for prev.Next.Val <= current.Val {
				prev = prev.Next
			}

			// current插入到排序链表中, 插入节点位置前都是有序的，插入位置后是无续的
			unsorted.Next = current.Next
			current.Next = prev.Next
			prev.Next = current
		}
		fmt.Printf("unsorted: ")
		PrintList(unsorted)
		current = unsorted.Next
	}
	return result.Next
}
