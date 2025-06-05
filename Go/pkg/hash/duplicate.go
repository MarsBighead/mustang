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
