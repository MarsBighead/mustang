package graph

// 399. 除法求值
// https://leetcode.cn/problems/evaluate-division/description/?envType=study-plan-v2&envId=top-interview-150
// ????
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hash := map[string]int{}
	// 方程组中每个变量编号
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := hash[a]; !ok {
			hash[a] = len(hash)
		}
		if _, ok := hash[b]; !ok {
			hash[b] = len(hash)
		}
	}
	n := len(hash)
	fathers, weights := make([]int, n), make([]float64, n)
	for i := range fathers {
		fathers[i] = i
		weights[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fathers[x] != x {
			f := find(fathers[x])
			weights[x] *= weights[fathers[x]]
			fathers[x] = f
		}
		return fathers[x]
	}
	merge := func(from, to int, val float64) {
		f, t := find(from), find(to)
		weights[f] = val * weights[to] / weights[from]
		fathers[f] = t
	}
	for i, eq := range equations {
		merge(hash[eq[0]], hash[eq[1]], values[i])
	}
	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, s := hash[q[0]]
		end, e := hash[q[1]]
		if s && e && find(start) == find(end) {
			ans[i] = weights[start] / weights[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
