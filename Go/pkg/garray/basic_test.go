package garray

import (
	"fmt"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	nums := [][]int{
		{1, 7, 11},
		{2, 4, 6},
	}
	k := 3
	ans := kSmallestPairs(nums[0], nums[1], k)
	for _, p := range ans {
		fmt.Println(p)
	}
}
