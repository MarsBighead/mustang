package garray

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 双指针逼近
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	if (target > 0 && numbers[i] > target) || (target < 0 && target > numbers[j]) {
		return []int{}
	}
	for i < j {
		v := numbers[i] + numbers[j]
		if v == target {
			return []int{i + 1, j + 1}
		} else if v < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

// https://leetcode.cn/problems/sort-colors/
// 颜色排序，双指针
func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	for i, _ := range nums {
		if i < red {
			nums[i] = 0
		} else if i < red+white {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}

}
