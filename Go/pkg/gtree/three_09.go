package gtree

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
// 958. 二叉树的完全性检验
// 一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，
// 如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，
// 则这棵二叉树称为完全二叉树
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var (
		recursive func(*TreeNode, int)
		size      = 0
		maxIdx    = 0
	)
	recursive = func(root *TreeNode, index int) {

		if root == nil {
			return
		}
		size += 1
		if maxIdx < index {
			maxIdx = index
		}

		recursive(root.Left, index*2)

		recursive(root.Right, index*2+1)

	}
	recursive(root, 1)
	return size == maxIdx
}
