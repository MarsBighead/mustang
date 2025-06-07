package gdp

/*
gdp: go dynamics programming
*/
//https://leetcode.cn/problems/n-th-tribonacci-number/description/?envType=study-plan-v2&envId=dynamic-programming
// 1137. 第 N 个泰波那契数
func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	t0, t1, t2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		//t :=
		t0, t1, t2 = t1, t2, t0+t1+t2
	}
	return t2

}
