package graph

// 130. 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/description/?envType=study-plan-v2&envId=top-interview-150
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'Z'

		for _, pos := range [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
			dfs(pos[0], pos[1])
		}
	}
	// 遍历左右
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	// 遍历上下
	for i := 1; i < n-1; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}
		}
	}
}

// https://leetcode.cn/problems/clone-graph/description/?envType=study-plan-v2&envId=top-interview-150
// 133. 克隆图
// 哈希表存储所有已被访问和克隆的节点
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var deepClone func(node *Node) *Node
	deepClone = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		clonedNode := &Node{node.Val, []*Node{}}
		visited[node] = clonedNode
		for _, n := range node.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, deepClone(n))
		}
		return clonedNode
	}

	return deepClone(node)
}
