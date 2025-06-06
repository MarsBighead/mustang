package garray

// 852. 山脉数组的峰顶索引
// https://leetcode.cn/problems/peak-index-in-a-mountain-array/
func peakIndexInMountainArray(arr []int) int {
	peak, length := 0, len(arr)
	for i, v := range arr {
		if v > arr[peak] {
			peak = i
			if i < length-1 && v > arr[i+1] {
				break
			}
		}
	}
	return peak
}
