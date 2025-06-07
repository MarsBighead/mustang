package linked

// https://leetcode.cn/problems/odd-even-linked-list/
// 328. 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	// 偶数
	var (
		even = new(ListNode)
		i    int
	)
	prev := head
	current := head
	next := even
	for current != nil {
		i += 1
		if i%2 == 0 {
			prev.Next = current.Next
			next.Next = &ListNode{Val: current.Val}
			next = next.Next
		} else {
			prev = current
		}
		current = current.Next
	}
	if even.Next != nil {
		prev.Next = even.Next
	}
	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
