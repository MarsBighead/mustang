package gtree

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/description/?envType=study-plan-v2&envId=top-interview-150
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Right)
		if node.Val > ans {
			ans = node.Val
		}
		traversal(node.Left)
	}
	traversal(root)
	return ans
}

// https://leetcode.cn/problems/invert-binary-tree/description/
// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 230. 二叉搜索树中第 K 小的元素
//https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历变形
func kthSmallest(root *TreeNode, k int) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	ans := 0
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil || k == 0 {
			return
		}
		traversal(node.Left)
		k--
		if k == 0 {
			ans = node.Val // 根
		}
		traversal(node.Right)
	}
	traversal(root)
	return ans
}

// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if right == nil {
		return left
	}
	return right

}
