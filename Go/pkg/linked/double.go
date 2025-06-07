package linked

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
