package gtree

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
// 637. 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	ans := [][2]int{}
	var dfs func(level int, node *TreeNode)
	dfs = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		if level < len(ans) {
			ans[level][0] += node.Val
			ans[level][1]++
		} else {
			ans = append(ans, [2]int{node.Val, 1})
		}
		dfs(level+1, node.Left)
		dfs(level+1, node.Right)
	}
	dfs(0, root)
	averages := make([]float64, len(ans))
	for i, data := range ans {
		averages[i] = float64(data[0]) / float64(data[1])
	}
	return averages
}
