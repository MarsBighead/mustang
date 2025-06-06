package number

// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	hash := make(map[int]bool, n)

	for _, num := range nums {
		hash[num] = true
	}
	if len(hash) == 1 {
		return 1
	}
	ans := 0
	for num, _ := range hash {
		if !hash[num-1] {
			current := num
			cnt := 1
			for hash[current+1] {
				current++
				cnt++
			}
			if ans < cnt {
				ans = cnt
			}
		}

	}
	return ans
}
