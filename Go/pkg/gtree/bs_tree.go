package gtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []int
	index int
}

// 173. 二叉搜索树迭代器
// https://leetcode.cn/problems/binary-search-tree-iterator/
func Constructor(root *TreeNode) BSTIterator {
	var (
		traversal func(*TreeNode)
	)
	bst := BSTIterator{index: -1}
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			traversal(node.Left)
		}
		bst.stack = append(bst.stack, node.Val)
		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root)
	return bst
}

func (this *BSTIterator) Next() int {
	this.index += 1
	return this.stack[this.index]
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.stack)-1
}

// 700. 二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val == val {
		return root
	}
	return nil
}

func searchBSTv1(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode
	for root != nil {
		if root.Val == val {
			ans = root
			break
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return ans
}
