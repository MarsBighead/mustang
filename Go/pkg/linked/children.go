package linked

// 430. 扁平化多级双向链表
// https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/description/
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Next == nil && root.Child == nil {
		return root
	}

	var (
		current = root
		stack   []*Node
		n       = len(stack)
	)
	for current.Next != nil || current.Child != nil || n > 0 {
		if current.Child != nil {
			if current.Next != nil {
				current.Next.Prev = nil
				stack = append(stack, current.Next)
				n++
			}
			current.Next = current.Child
			current.Child.Prev = current
			current.Child = nil
		}
		if current.Next == nil && n > 0 {
			current.Next = stack[len(stack)-1]
			stack[len(stack)-1].Prev = current
			stack = stack[:len(stack)-1]
			n--
		}
		current = current.Next
	}
	return root
}
