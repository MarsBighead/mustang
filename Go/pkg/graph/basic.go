package graph

// 200.岛屿数量
// https://leetcode.cn/problems/number-of-islands/description/?envType=study-plan-v2&envId=top-interview-150
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > rows || y < 0 || y > columns {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		for _, pos := range [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
			dfs(pos[0], pos[1])
		}
	}
	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if grid[r][c] == '1' {
				ans++
				dfs(r, c)
			}
		}
	}
	return ans
}
