package gtree

import (
	"fmt"
)

// https://leetcode.cn/problems/same-tree/
// 100. 相同的树
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
		// a.Val != b.Val
		return false
	}

	return check(p, q)
}

// https://leetcode.cn/problems/symmetric-tree/description/
// 101. 对称二叉树
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

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
// 103. 二叉树的锯齿形层序遍历
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

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度
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

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTreev1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	idx := 0
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[idx] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[idx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	idx := len(inorder) - 1
	root := &TreeNode{Val: postorder[idx]}
	stack := []*TreeNode{root}
	for i := len(inorder) - 2; i >= 0; i-- {
		pval := postorder[i]
		node := stack[len(stack)-1]
		fmt.Printf("node: %d, pval: %d, inorder: %d\n", stack[len(stack)-1].Val, pval, inorder[idx])
		if node.Val != inorder[idx] {
			node.Right = &TreeNode{Val: pval}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx--
			}
			node.Left = &TreeNode{Val: pval}
			stack = append(stack, node.Left)
		}
	}
	return root
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	idx := 0
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		pval := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != postorder[idx] {
			node.Left = &TreeNode{pval, nil, nil}
			stack = append(stack, node.Left)
		} else {
			fmt.Printf("len stack: %d, node: %d, pval: %d, postorder: %d\n", len(stack), stack[len(stack)-1].Val, pval, postorder[idx])
			for len(stack) != 0 && stack[len(stack)-1].Val == postorder[idx] {
				node = stack[len(stack)-2]
				stack = stack[:len(stack)-1]
				idx++
			}
			node.Right = &TreeNode{pval, nil, nil}
			stack = append(stack, node.Right)
		}
	}

	return root
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
// 107. 二叉树的层序遍历 II
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

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 111. 二叉树的最小深度
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

// https://leetcode.cn/problems/path-sum/description/?envType=study-plan-v2&envId=top-interview-150
// 112. 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	nextSum := targetSum - root.Val
	return hasPathSum(root.Left, nextSum) || hasPathSum(root.Right, nextSum)
}

func hasPathSumv1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodes := []*TreeNode{root}
	values := []int{root.Val}
	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		tmp := values[0]
		values = values[1:]
		if node.Left == nil && node.Right == nil {
			if tmp == targetSum {
				return true
			}
			continue
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			values = append(values, node.Left.Val+tmp)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			values = append(values, node.Right.Val+tmp)
		}

	}
	return false
}

// https://leetcode.cn/problems/path-sum-ii/description/
// 113. 路径总和 II
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

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
// 124. 二叉树中的最大路径和
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

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/?envType=study-plan-v2&envId=top-interview-150
// 129. 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	ans := 0
	var traversal func(sum int, node *TreeNode)
	traversal = func(sum int, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			ans += node.Val + sum*10
		}

		//fmt.Println(sum, level)
		sum = node.Val + sum*10
		if node.Left != nil {
			traversal(sum, node.Left)
		}
		if node.Right != nil {
			traversal(sum, node.Right)
		}

	}
	traversal(0, root)
	return ans
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 144. 二叉树的前序遍历
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

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		list[i-1].Left, list[i-1].Right = nil, list[i]
	}

}

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
// 145. 二叉树的后序遍历
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

func (bst *BSTIterator) Next() int {
	bst.index += 1
	return bst.stack[bst.index]
}

func (bst *BSTIterator) HasNext() bool {
	return bst.index < len(bst.stack)-1
}

// https://leetcode.cn/problems/binary-tree-right-side-view/description/
// 199. 二叉树的右视图
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
