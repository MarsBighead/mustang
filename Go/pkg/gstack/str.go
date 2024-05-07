package gstack

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 未通过
func decodeString(s string) string {
	prefix := []rune{}
	suffix := []rune{}
	stack := []rune{}
	block, preblock := 0, 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			preblock = block
			block++
			stack = append(stack, ch)
		} else if ch == ']' {
			// 需要针Block，一个一个从内向外处理
			// n*str
			preblock = block
			block--
		} else if ch >= 'a' && ch <= 'z' {
			if block == 0 && preblock == 0 {
				prefix = append(prefix, ch)
			} else if block == 0 && preblock != 0 {
				suffix = append(suffix, ch)
			} else {
				stack = append(stack, ch)
			}
		}
	}
	fmt.Printf("stack=%s,suffix=%s\n", string(stack), string(suffix))
	//result := ""
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] >= '0' && stack[i] <= '9' {
			n := int(stack[i] - '0')
			ss := string(stack[i+1:])
			stack = stack[:i]
			ss = strings.Repeat(ss, n)
			stack = append(stack, []rune(ss)...)
			//fmt.Println("xxx=", string(stack), string(ss), n)
			i = len(stack) - 1
		}
	}

	return string(prefix) + string(stack) + string(suffix)
}

/*
数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n的最大数。
例如A={1, 2, 4, 9}，x=2533，返回2499
*/
func getLargerButLessThanK(nums []int, k int) int {
	ks := []int{k % 10}
	for k > 0 {
		k = k / 10
		if k > 0 {
			ks = append(ks, k%10)
		}
	}
	result := 0
	sort.Ints(nums)
	if len(nums) < len(ks) {
		for i := 0; i < len(nums); i++ {
			result += nums[i] * int(math.Pow10(i))
		}
	} else {
		length := len(ks)
		top := ks[length-1]
		for i := 0; i < length-1; i++ {
			if top > 10 {
				top = top + ks[length-i-1]*10
			} else {
				top = ks[length-i-1]
			}

			max := nums[0]
			for j := 0; j < len(nums); j++ {
				if nums[j] > top {
					nums = nums[j:]
					break
				} else {
					max = nums[j]
				}
			}
			fmt.Println(i, ks, max, top)
			if max > top {
				ks = nums[len(nums)-length+1:]
				break
			} else {
				if top > 10 {
					ks[length-i-1] = max
				} else {
					ks[length-i-1] = top
				}
			}
		}

		for i := 0; i < len(ks); i++ {
			result += ks[i] * int(math.Pow10(i))
		}

	}
	return result
}
