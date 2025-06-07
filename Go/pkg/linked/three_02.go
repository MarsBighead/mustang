package linked

// 203. 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	return reverse(cur, head)
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
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

func reverse(prev, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	} else {
		node := cur.Next
		cur.Next = prev
		return reverse(cur, node)
	}

}

// https://leetcode.cn/problems/palindrome-linked-list/
// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var (
		ok      bool
		prev    *ListNode
		current = head
		tail    = new(ListNode)
		raw     = tail
	)

	for current != nil {
		tail.Next = &ListNode{Val: current.Val}
		tail = tail.Next
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ok = true
	for prev.Next != nil && raw.Next.Next != nil {
		if prev.Val != raw.Next.Val {
			ok = false
			break
		}
		prev = prev.Next
		raw = raw.Next

	}

	return ok
}

// https://leetcode.cn/problems/palindrome-linked-list/
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var (
		array   []int
		current = head
		cnt     int
	)

	for current != nil {
		cnt += 1
		array = append(array, current.Val)
		current = current.Next
	}
	for i := 0; i < cnt/2; i++ {
		if array[i] != array[cnt-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome3(head *ListNode) bool {
	// 快慢指针
	slow, fast := head, head
	// 找到链表中点, slow移动一个节点，fast移动两个节点，因此fast完成时，slow指向链表后半段
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半部分链表
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	// 比较前后两部分链表
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
