package linked

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list-ii/description/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var (
		prev       = new(ListNode)
		start, end *ListNode
		i, j       int
	)
	current := head
	for current != nil {
		i += 1
		if i > right {
			end = current
			break
		}
		if i < left {
			start = current
		}
		next := current.Next
		if i >= left && i <= right {
			current.Next = prev
			prev = current
		}
		current = next
	}

	if start != nil {
		start.Next = prev
	} else {
		head = prev
	}

	current = head

	for current != nil {
		j += 1
		if j == right {
			current.Next = end
			break
		}
		current = current.Next
	}

	return head
}

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
	if head == nil {
		fmt.Println()
		return
	}
	fmt.Print(head.Val, " ")
	for head.Next != nil {
		head = head.Next
		fmt.Print(head.Val, " ")
	}
	fmt.Println()
	return
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node.Next != nil {
		fmt.Println(node.Val)
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}

		if node.Next == nil {
			break
		}

	}
	return head
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
func deleteAllDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 创建一个虚拟头结点
	dummy := &ListNode{Next: head}
	next := head
	current := dummy
	for next != nil {
		flag := false
		// 当前节点值
		val := next.Val
		// 下一个节点

		// 如果下一个节点存在并且值与当前节点值相同，则一直移动next直到值不同
		for next.Next != nil && val == next.Next.Val {
			next = next.Next
			flag = true
		}
		// 如果next为空，表示所有重复节点都被跳过了，直接链接
		if flag {
			current.Next = next.Next
		} else {
			current.Next = next
			current = next
		}
		// 移动到下一个节点
		next = next.Next
	}
	return dummy.Next
}

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
