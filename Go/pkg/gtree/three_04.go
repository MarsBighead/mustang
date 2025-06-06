package gtree

import "fmt"

// 450. 删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/description/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	var (
		parent *TreeNode
	)
	for node != nil && node.Val != key {
		parent = node
		if node.Val > key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if node == nil {
		return root
	}
	if node.Left == nil && node.Right == nil {
		node = nil
	} else if node.Right == nil {
		node = node.Left
	} else if node.Left == nil {
		node = node.Right
	} else {
		children, childrenParent := node.Right, node
		// 把右子树最小值，替换发现key的node节点位置，其parent对应该节点置空
		// 把原来的子树位置挂到children左右子树位置下
		for children.Left != nil {
			childrenParent = children
			children = children.Left
		}
		fmt.Printf("children: %d, children parent:%d, node: %d \n", children.Val, childrenParent.Val, node.Val)
		if childrenParent.Val == node.Val {
			childrenParent.Right = children.Right
		} else {
			childrenParent.Left = children.Right
		}
		fmt.Printf("XX: children: %d, children parent:%d, left: %#v\n", children.Val, childrenParent.Val, childrenParent.Left)
		children.Right = node.Right
		children.Left = node.Left
		node = children
	}

	if parent == nil {
		return node
	}
	if parent.Left != nil && parent.Left.Val == key {
		parent.Left = node
	} else {
		parent.Right = node
	}

	return root
}
