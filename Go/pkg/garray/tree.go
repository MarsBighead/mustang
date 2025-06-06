package garray

import (
	"encoding/json"
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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
func inorderTraversalv1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, output := []*TreeNode{}, []int{}

	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			l := len(stack)
			node := stack[l-1]
			stack = stack[:l-1]
			output = append(output, node.Val)
			root = node.Right
		}
	}

	return output
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var (
		nodes     [][]*Node
		recursive func(*Node, int)
	)
	//nodes = append(nodes, []*Node{root})
	recursive = func(node *Node, i int) {
		//length := len(nodes)
		if node == nil {
			return
		}
		if len(nodes) < i {
			nodes = append(nodes, []*Node{node})
		} else {
			length := len(nodes[i-1])
			nodes[i-1][length-1].Next = node
			nodes[i-1] = append(nodes[i-1], node)
		}
		i++
		recursive(node.Left, i)
		recursive(node.Right, i)
	}
	recursive(root, 1)

	return root
}

func connectv2(root *Node) *Node {
	if root == nil {
		return root
	}
	var (
		nodes     []*Node
		recursive func(*Node, int)
	)
	//nodes = append(nodes, []*Node{root})
	recursive = func(node *Node, i int) {
		//length := len(nodes)
		if node == nil {
			return
		}
		if len(nodes) < i {
			nodes = append(nodes, node)
		} else {
			nodes[i-1].Next = node
			nodes[i-1] = node
		}
		i++
		recursive(node.Left, i)
		recursive(node.Right, i)
	}
	recursive(root, 1)

	return root
}

func parseBrConifgOptions(body string) ([]string, error) {
	var options []string
	err := json.Unmarshal([]byte(body), &options)
	if err != nil {
		return nil, err
	}
	return options, nil

}
