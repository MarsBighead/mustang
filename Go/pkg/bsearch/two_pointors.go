package bsearch

import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	var (
		i, j       int
		slow, fast = head, head
		prev       *ListNode
	)

	for fast != nil && fast.Next != nil {
		i += 1
		j += 2
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil && fast.Next == nil {
		j += 1
	}
	pos := j - n + 1
	var (
		current = slow
		result  = &ListNode{Next: head}
		x       = i
	)
	if pos <= i {
		current = head
		prev = result
		x = 0
	}
	fmt.Printf("Half: %d, end: %d, x: %d, pos: %d\n", i, j, x, pos)
	PrintList(slow)
	for current != nil {
		x += 1
		if x == pos {
			PrintList(prev)
			prev.Next = current.Next
		}
		prev = current
		current = current.Next
	}
	//PrintList(head)
	return result.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	var (
		dummy      = &ListNode{Next: head}
		slow, fast = dummy, dummy
	)

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next == head {
		return head
	}
	hash := make(map[*ListNode]struct{})
	current := head
	for current != nil {
		if _, ok := hash[current]; ok {
			return current
		} else {
			hash[current] = struct{}{}
		}
		current = current.Next
	}

	return nil

}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next == head {
		return head
	}
	var (
		slow, fast = head, head
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			node := head
			for node != slow {
				node = node.Next
				slow = slow.Next
			}
			return node

		}

	}

	return nil

}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}
	var (
		a    = headA
		b    = headB
		x, y int
	)
	for a != nil {
		x += 1
		a = a.Next
	}
	for b != nil {
		y += 1
		b = b.Next
	}
	a = headA
	b = headB
	if x < y {
		a, b = b, a
		x, y = y, x
	}
	gap := x - y
	for gap > 0 {
		a = a.Next
	}
	if a == b {
		return a
	}
	for a != b {
		a = a.Next
		b = b.Next
		if a == b {
			return a
		}
	}
	return nil
}

// https://leetcode.cn/problems/reorder-list/description/
func reorderList(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil

	fast = reverse(fast)
	slow = head
	PrintList(fast)
	PrintList(head)
	for fast != nil && slow != nil {
		tmp := slow.Next
		slow.Next = fast
		slow = tmp
		tmp = fast.Next
		fast.Next = slow
		fast = tmp
	}
	PrintList(head)

}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var (
		prev    *ListNode
		current = head
	)
	for current != nil {
		// 真正列表的next内容
		next := current.Next

		// 反转next指向prev
		current.Next = prev
		// prev反转增加current节点
		prev = current

		// current指向下一个节点
		current = next
	}
	return prev
}

// https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		tail, head *ListNode
		carry      int
	)
	for l1 != nil || l2 != nil {
		var (
			n, m int
		)
		if l1 != nil {
			n = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m = l2.Val
			l2 = l2.Next
		}
		sum := n + m + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}

func addTwoNumbersv2(l1 *ListNode, l2 *ListNode) *ListNode {
	z, len, carry := 0, 0, 0
	var (
		x, y []*ListNode
	)
	n, m := l1, l2
	for n != nil {
		z += 1
		x = append(x, n)
		n = n.Next
	}
	len = z
	for m != nil {
		z -= 1
		y = append(y, m)
		m = m.Next
	}
	fmt.Println("short:", z, len)
	if z < 0 {
		l1 = l2
		x, y = y, x
		len, z = len-z, len
	} else {
		z = len - z
	}
	result := &ListNode{Next: l1}

	for i := 0; i < len; i++ {
		sum := x[len-i-1].Val + carry
		if z-i-1 >= 0 {
			sum = sum + y[z-i-1].Val
		}
		sum, carry = sum%10, sum/10
		x[len-i-1].Val = sum
	}

	if carry > 0 {
		result.Val = carry
		return result
	}

	return result.Next
}
