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
