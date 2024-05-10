package garray

import "fmt"

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
