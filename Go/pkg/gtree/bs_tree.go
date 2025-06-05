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

// 701. 二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/
// 该解答未实现插入值，替换原来节点，原有节点移动重排的情况
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				break
			}
			node = node.Left
		} else if node.Val < val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				break
			}
			node = node.Right
		}
	}
	return root
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
