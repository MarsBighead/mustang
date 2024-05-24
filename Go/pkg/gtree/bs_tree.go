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
