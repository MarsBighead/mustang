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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
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
