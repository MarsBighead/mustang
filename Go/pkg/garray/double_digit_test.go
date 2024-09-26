package garray

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	ans := CombinationSum(candidates, target)
	fmt.Printf("%#v\n", ans)
}

func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ans := CombinationSum2(candidates, target)
	fmt.Printf("%#v\n", ans)
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	n := FirstMissingPositive(nums)
	fmt.Printf("output %#v\n", n)
	nums = []int{3, 4, -1, 1}
	n = FirstMissingPositiveV1(nums)
	fmt.Printf("output %#v\n", n)
}
