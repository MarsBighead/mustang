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

// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

func preorderTraversalv1l(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	stack, output := []*TreeNode{root}, []*TreeNode{}

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		if node != nil {
			output = append(output, node)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}

	}

	return output
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

func postorderTraversalv1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		output []int
		prev   *TreeNode
		stack  []*TreeNode
	)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			output = append(output, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
		fmt.Println("xxx", output)
	}

	return output

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

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//可以不贡献路径
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)

	}
	return check(root, root)
}

// https://leetcode.cn/problems/path-sum-ii/description/
// 路径总和II s s
func pathSum(root *TreeNode, targetSum int) [][]int {

	path := []int{}
	var (
		dfs func(*TreeNode, int)
		ans = make([][]int, 0)
	)
	// left represents for leave
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return ans
}

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

// https://leetcode.cn/problems/binary-tree-right-side-view/description/
// 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		level func(*TreeNode, int)
	)
	ans := []int{}
	n := 1
	level = func(node *TreeNode, n int) {
		if len(ans) < n {
			ans = append(ans, node.Val)
		}
		ans[n-1] = node.Val
		n += 1
		if node.Left != nil {
			level(node.Left, n)
		}
		if node.Right != nil {
			level(node.Right, n)
		}
	}
	level(root, n)
	return ans

}

// https://leetcode.cn/problems/invert-binary-tree/description/
// 反转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// https://leetcode.cn/problems/subtree-of-another-tree/description/
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

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
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

// https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
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

	return check(p, q)
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
