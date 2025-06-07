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
		if x < 0 || x >= rows || y < 0 || y >= columns {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
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

// https://leetcode.cn/problems/course-schedule/description/?envType=study-plan-v2&envId=top-interview-150
// 207. 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		ok      = true
		dfs     func(u int)
	)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	// 未搜索: 0，搜索中: 1, 搜索完成:2
	dfs = func(x int) {
		visited[x] = 1
		for _, y := range edges[x] {
			if visited[y] == 0 {
				dfs(y)
				if !ok {
					return
				}
			} else if visited[y] == 1 {
				ok = false
				return
			}
		}
		visited[x] = 2
		result = append(result, x)
	}

	for i := 0; i < numCourses && ok; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if !ok {
			return ok
		}
	}
	return ok

}

// https://leetcode.cn/problems/course-schedule-ii/description/?envType=study-plan-v2&envId=top-interview-150
// 210. 课程表 II
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		ok      = true
		dfs     func(u int)
	)
	for _, info := range prerequisites {
		// edges: 存储前置课程列表
		edges[info[0]] = append(edges[info[0]], info[1])
	}
	// 未搜索: 0，搜索中: 1, 搜索完成:2
	dfs = func(x int) {
		visited[x] = 1
		for _, y := range edges[x] {
			if visited[y] == 0 {
				dfs(y)
				if !ok {
					return
				}
			} else if visited[y] == 1 {
				ok = false
				return
			}
		}
		visited[x] = 2
		result = append(result, x)
	}

	for i := 0; i < numCourses && ok; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if !ok {
			return []int{}
		}
	}
	return result

}
