package linked

import "fmt"

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/?envType=study-plan-v2&envId=top-interview-150
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	current := head
	var (
		start = head
		end   *ListNode
	)
	root := &ListNode{Next: start}
	n := 0
	for current != nil {
		n++
		if n%k == 0 {
			// current就是k段最后节点
			// end是上次翻转的最后节点，也就是上次未翻转前的start
			var prev *ListNode
			node := start
			for i := 0; i < k; i++ {
				// 真正列表的next内容
				next := node.Next
				// 反转next指向prev
				node.Next = prev
				prev = node
				// current指向下一个节点
				node = next
			}
			if n == k {
				root.Next = prev
			} else {
				end.Next = prev
			}
			end = start
			start.Next = node
			current = node
			start = node
			PrintList(root.Next)
		} else {
			current = current.Next
		}
	}
	return root.Next
}

// https://leetcode.cn/problems/rotate-list/description/
// 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	var (
		length, i int
		current   = head
		rotate    *ListNode
		tail      *ListNode
	)
	for current != nil {
		length += 1
		if current.Next == nil {
			tail = current
		}
		current = current.Next
	}
	//tail = current
	r := k % length
	if r == 0 {
		return head
	}
	current = head
	for current != nil {
		i += 1
		if i == length-r {
			rotate = current.Next
			current.Next = nil
			break
		}
		current = current.Next
	}
	tail.Next = head
	return rotate
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
// 80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	left := 2
	for i := 2; i < n; i++ {
		if nums[left-2] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

// 82. 删除排序链表中的重复元素 II
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

// 83. 删除排序链表中的重复元素
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

// https://leetcode.cn/problems/reverse-linked-list-ii/description/
// 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 你可以使用一趟扫描完成反转吗？
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
