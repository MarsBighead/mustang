package number

// https://leetcode.cn/problems/palindrome-number
// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	//fmt.Println(x, reversed)
	return x == reversed || x == reversed/10

}

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
