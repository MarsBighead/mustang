package gtree

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/?envType=study-plan-v2&envId=top-interview-150
// 530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	if root.Left != nil && root.Right == nil {
		return root.Val - root.Left.Val
	}
	if root.Left == nil && root.Right != nil {
		return root.Right.Val - root.Val
	}
	ans, pre := 100000, -1
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		tmp := n.Val - pre
		if pre != -1 && tmp < ans {
			ans = tmp
		}
		pre = n.Val
		dfs(n.Right)
	}

	dfs(root)
	return ans
}

// https://leetcode.cn/problems/subtree-of-another-tree/description/
// 572. 另一棵树的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}
	var check func(*TreeNode, *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val == b.Val {
			return check(a.Left, b.Left) && check(a.Right, b.Right)
		}
		return false
	}
	if check(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
