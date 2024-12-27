package garray

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	mx := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	nums := spiralOrder(mx)
	fmt.Println(nums)
}

func TestIsPossibleDivide(t *testing.T) {
	ks := []int{4}
	nss := [][]int{{1, 2, 3, 3, 4, 4, 5, 6}}
	for i, nums := range nss {
		k := ks[i]
		ok := isPossibleDivide(nums, k)
		fmt.Println("ok=", ok)
	}
}
