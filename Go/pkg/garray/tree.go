package garray

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Construct(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	tree := constructTree(nums, 0)
	return tree

}

func constructTree(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}
	return &TreeNode{
		Val:   nums[index],
		Left:  constructTree(nums, 2*index+1),
		Right: constructTree(nums, 2*index+2),
	}
}

// 左序优先遍历
func printTree(root *TreeNode) {
	//fmt.Printf("Print: %#v\n", root)
	if root == nil {
		fmt.Println("nil")
		return
	}
	fmt.Print(root.Val, " ")
	printTree(root.Left)
	printTree(root.Right)
	fmt.Println()
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		list[i-1].Left, list[i-1].Right = nil, list[i]
	}
	return

}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func preorderTraversalv2(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var (
		vals     []int
		preorder func(*TreeNode)
	)

	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return vals
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

func inorderTraversalv2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
// 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//fmt.Println("length of head:", len(root))
	var (
		result    []int
		postorder func(*TreeNode)
	)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}
	postorder(root)

	return result

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result = make([][]int, 0)
		level  func(*TreeNode, int)
		n      = 1
	)
	level = func(node *TreeNode, n int) {
		if len(result) < n {
			tmp := make([]int, 0)
			result = append(result, tmp)
		}
		result[n-1] = append(result[n-1], node.Val)
		n += 1
		if node.Left != nil {
			level(node.Left, n)
		}
		if node.Right != nil {
			level(node.Right, n)
		}
	}
	level(root, n)
	return result
}

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
// 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result      = make([][]int, 0)
		zigzagLevel func(*TreeNode, int)
		n           = 1
	)
	zigzagLevel = func(node *TreeNode, n int) {
		if len(result) < n {
			tmp := make([]int, 0)
			result = append(result, tmp)
		}
		result[n-1] = append(result[n-1], node.Val)
		n += 1

		if node.Left != nil {
			zigzagLevel(node.Left, n)
		}
		if node.Right != nil {
			zigzagLevel(node.Right, n)
		}

	}
	zigzagLevel(root, n)
	for i := 0; i < len(result); i++ {
		if i%2 == 1 {
			for j, n := 0, len(result[i]); j < n/2; j++ {
				result[i][j], result[i][n-1-j] = result[i][n-1-j], result[i][j]
			}
		}
	}
	return result
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
// 二叉树 自底向上的层序遍历
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result     = make([][]int, 0)
		levelOrder func(*TreeNode, int)
		n          = 1
	)
	levelOrder = func(node *TreeNode, n int) {
		if len(result) < n {
			tmp := make([]int, 0)
			result = append(result, tmp)
		}
		result[n-1] = append(result[n-1], node.Val)
		n += 1

		if node.Left != nil {
			levelOrder(node.Left, n)
		}
		if node.Right != nil {
			levelOrder(node.Right, n)
		}

	}
	levelOrder(root, n)
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	result := 1 << 31
	if root.Left != nil {
		result = min(minDepth(root.Left), result)
	}
	if root.Right != nil {
		result = min(minDepth(root.Right), result)
	}
	return result + 1

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
