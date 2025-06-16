package gtree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// https://leetcode.cn/problems/construct-quad-tree/description/?envType=study-plan-v2&envId=top-interview-150
// 427. 建立四叉树
func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}
	var build func(grid [][]int, start, end int) *Node
	build = func(grid [][]int, start, end int) *Node {
		for _, row := range grid {
			for _, v := range row[start:end] {
				if v != grid[0][start] { // 不是叶节点
					m, n := len(grid)/2, (start+end)/2
					return &Node{
						true,
						false,
						build(grid[:m], start, n),
						build(grid[:m], n, end),
						build(grid[m:], start, n),
						build(grid[m:], n, end),
					}
				}
			}
		}
		return &Node{Val: grid[0][start] == 1, IsLeaf: true}
	}
	return build(grid, 0, len(grid))
}
