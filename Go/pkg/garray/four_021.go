package garray

// https://leetcode.cn/problems/divide-array-in-sets-of-k-consecutive-numbers/description/
// 1296. 划分数组为连续数字的集合
func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	for _, x := range nums {
		if hash[x] == 0 {
			continue
		}
		for num := x; num < x+k; num++ {
			if hash[num] == 0 {
				return false
			}
			hash[num]--
		}
	}

	//fmt.Printf("%#v\n%#v\n", nums, hash)
	return true
}
