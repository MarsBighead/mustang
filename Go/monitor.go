package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println("before sorted", nums)
	nums = []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))

	nums = []int{2, 2, 1, 3, 3}
	fmt.Println("single numbers:", singleNumber(nums))
	/*	insertSort(nums)
		fmt.Println("insert sorted", nums)
		selectionSort(nums)
		fmt.Println("selection sorted", nums)

		nums[4] = 6
		bubbleSort(nums)
		fmt.Println("sorted", nums)
	*/
	/*
		ctx, cancel := context.WithCancel(context.Background())
		go watch(ctx, "【监控1】")
		go watch(ctx, "【监控2】")
		go watch(ctx, "【监控3】")
		time.Sleep(10 * time.Second)
		fmt.Println("可以了，通知监控停止")
		//为了检测监控过是否停止，如果没有监控输出，就表示停止了
		cancel()
		time.Sleep(2 * time.Second)
	*/
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

// rotate: https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	length := len(nums)
	r := k % length

	var tmp []int
	tmp = append(tmp, nums[length-r:length]...)
	for i := 0; i < length-r; i++ {
		nums[length-1-i] = nums[length-1-i-r]
	}
	for i := 0; i < r; i++ {
		nums[i] = tmp[i]
	}
}

func bubbleSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			flag = true
			if !flag {
				break
			}
		}
	}
	return

}

func selectionSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		mini := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[mini] {
				mini = j
				//nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}
		if i != mini {
			nums[i], nums[mini] = nums[mini], nums[i]
		}
	}
	return
}

func insertSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		tmp := nums[i]
		j := i
		for j > 0 && nums[j-1] > tmp {
			nums[j] = nums[j-1]
			j = j - 1
		}
		nums[j] = tmp
		fmt.Println(nums)
		//break
	}
	return
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	hash := make(map[int]int)
	isDup := false
	for _, n := range nums {

		if v, ok := hash[n]; ok {
			hash[n] = v + 1
			isDup = true
			break
		} else {
			hash[n] = 1
		}
	}
	return isDup
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) <= 1 {
		return false
	}
	isDup := false
	hash := make(map[int]int)
	for i, n := range nums {
		if j, ok := hash[n]; ok {
			if i-j <= indexDiff && nums[i]-nums[j] <= valueDiff {
				isDup = true
				break
			}
		} else {
			hash[n] = i
		}
	}

	return isDup

}

// bitMap
func singleNumber(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}

	return x

}
