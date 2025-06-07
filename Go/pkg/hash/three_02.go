package hash

import "strings"

// 202. 快乐数
// https://leetcode.cn/problems/happy-number/description/
func isHappy(n int) bool {
	sumSquare := func(num int) int {
		sum := 0
		for num > 0 {
			x := num % 10
			sum += x * x
			num = num / 10
		}
		return sum
	}
	slow, fast := n, sumSquare(n)
	for slow != fast {
		slow = sumSquare(slow)
		fast = sumSquare(fast)
		fast = sumSquare(fast)
		//fmt.Println(slow, fast, cnt)

	}
	return slow == 1
}

func isHappyv1(n int) bool {
	sumSquare := func(num int) int {
		sum := 0
		for num > 0 {
			x := num % 10
			sum += x * x
			num = num / 10
		}
		return sum
	}
	hash := make(map[int]int)

	for {
		n = sumSquare(n)
		hash[n]++
		if n == 1 || hash[n] > 1 {
			break
		}
	}
	return n == 1
}

// https://leetcode.cn/problems/contains-duplicate-ii/description
// 219. 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k == 0 {
		return false
	}
	hash := make(map[int]int)
	for i, num := range nums {
		if p, ok := hash[num]; ok && i-p <= k {
			return true
		} else {
			hash[num] = i
		}
	}

	return false

}

// https://leetcode.cn/problems/word-pattern/description/
// 290. 单词规律
func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}
	hash := make(map[byte]string)
	phash := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {

		m, n := pattern[i], ss[i]
		if _, ok := hash[m]; !ok {
			hash[m] = n
		}
		if _, ok := phash[n]; !ok {
			phash[n] = m
		}
		x, y := phash[n], hash[m]
		if x != m || y != n {
			return false
		}

	}
	return true
}
