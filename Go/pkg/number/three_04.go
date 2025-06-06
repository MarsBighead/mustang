package number

// 485. 最大连续 1 的个数
// https://leetcode.cn/problems/max-consecutive-ones/description/

func findMaxConsecutiveOnes(nums []int) int {
	var (
		max, cur int
	)
	for _, v := range nums {
		// 利用位运算0，来取得当前累加值，也就是count ones
		cur = (cur + 1) * v
		//fmt.Println(v, cur)
		if cur > max {
			max = cur
		}
	}
	return max
}
