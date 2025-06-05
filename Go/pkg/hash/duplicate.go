package hash

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/description/
// 位图算法
func findDuplicates(nums []int) []int {
	length := len(nums)
	count := make([]int, length)
	for i := 0; i < length; i++ {
		v := nums[i]
		count[v-1] += 1
	}
	nums = nums[:0]
	for i, v := range count {
		if v == 2 {
			nums = append(nums, i+1)
		}
	}
	return nums
}

func findDuplicatesv2(nums []int) []int {

	mask := 1 << 61
	n := 15
	mask = mask ^ n
	//fmt.Println(mask << 61)
	return nums
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

// https://leetcode.cn/problems/ransom-note/description
// 383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	m, n := len(ransomNote), len(magazine)
	if m > n {
		return false
	}
	hash := make(map[rune]int, 26)
	for _, e := range ransomNote {
		hash[e] += 1
	}
	for _, e := range magazine {
		if _, ok := hash[e]; ok {
			hash[e]--
		}
	}
	for _, val := range hash {
		if val > 0 {
			return false
		}
	}
	return false
}

func canConstructv1(ransomNote string, magazine string) bool {
	m, n := len(ransomNote), len(magazine)
	if m > n {
		return false
	}
	check := [26]int{}
	for _, e := range ransomNote {
		check[e-'a']++
	}
	for _, e := range magazine {
		check[e-'a']--
	}
	for _, val := range check {
		if val > 0 {
			return false
		}
	}
	return true
}
