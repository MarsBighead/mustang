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

func TestValidIPAddress(t *testing.T) {
	addresses := []string{
		"172.16.254.1",
		"0.0.0.0",
		"1.0.01.0",
		"0db8",
		"1e1.4.5.6",
		"2001:0db8:85a3:0:0:8A2E:0370:7334:",
		"2001:0db8:85a3:0:0:8A2E:0370:7334",
		"2001:0db8:85a3:00000:0:8A2E:0370:7334",
	}
	for _, addr := range addresses {
		r := ValidIPAddress(addr)
		fmt.Println(r)
	}
}
