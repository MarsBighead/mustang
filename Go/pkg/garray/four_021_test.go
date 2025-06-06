package garray

import (
	"fmt"
	"testing"
)

func TestIsPossibleDivide(t *testing.T) {
	ks := []int{4}
	nss := [][]int{{1, 2, 3, 3, 4, 4, 5, 6}}
	for i, nums := range nss {
		k := ks[i]
		ok := isPossibleDivide(nums, k)
		fmt.Println("ok=", ok)
	}
}
