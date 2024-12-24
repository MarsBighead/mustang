package garray

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	numbers := [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}}
	for _, nums := range numbers {
		NextPermutation(nums)
		fmt.Println(nums)
	}
}
