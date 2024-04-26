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
	fmt.Println(nums)
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	cancel()
	time.Sleep(2 * time.Second)
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
