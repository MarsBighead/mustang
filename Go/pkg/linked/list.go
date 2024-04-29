package linked

// https://leetcode.cn/problems/odd-even-linked-list/
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

// https://leetcode.cn/problems/palindrome-linked-list/
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
