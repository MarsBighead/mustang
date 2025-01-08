package gtree

import "fmt"

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTreev1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	idx := 0
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[idx] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[idx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	idx := len(inorder) - 1
	root := &TreeNode{Val: postorder[idx]}
	stack := []*TreeNode{root}
	for i := len(inorder) - 2; i >= 0; i-- {
		pval := postorder[i]
		node := stack[len(stack)-1]
		fmt.Printf("node: %d, pval: %d, inorder: %d\n", stack[len(stack)-1].Val, pval, inorder[idx])
		if node.Val != inorder[idx] {
			node.Right = &TreeNode{Val: pval}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx--
			}
			node.Left = &TreeNode{Val: pval}
			stack = append(stack, node.Left)
		}
	}
	return root
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	idx := 0
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		pval := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != postorder[idx] {
			node.Left = &TreeNode{pval, nil, nil}
			stack = append(stack, node.Left)
		} else {
			fmt.Printf("len stack: %d, node: %d, pval: %d, postorder: %d\n", len(stack), stack[len(stack)-1].Val, pval, postorder[idx])
			for len(stack) != 0 && stack[len(stack)-1].Val == postorder[idx] {
				node = stack[len(stack)-2]
				stack = stack[:len(stack)-1]
				idx++
			}
			node.Right = &TreeNode{pval, nil, nil}
			stack = append(stack, node.Right)
		}
	}

	return root
}

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, ans := []*TreeNode{}, []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			root = node.Right
		}
	}
	return ans
}
