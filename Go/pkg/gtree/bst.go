package gtree

import "fmt"

// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	min := 0 - 1<<31
	//fmt.Printf("min=%b\n", min)
	n := 0
	for len(stack) > 0 || root != nil {
		n++
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		fmt.Printf("loop: %d, min=%d, %#v\n", len(stack), min, stack[0])
		if len(stack) > 1 {
			fmt.Printf("stack[1]=%#v\n\n", stack[1])

		}
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
		fmt.Printf("root=%#v\n\n", root)
	}
	return true
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	stack, output := []*TreeNode{root}, []int{}

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		if node != nil {
			output = append(output, node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}

	}

	return output
}
