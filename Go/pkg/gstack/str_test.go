package gstack

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	s = "abc3[cd]xyz"
	s = decodeString(s)
	fmt.Println(s)
}
func TestGetLargerButLessThanK(t *testing.T) {

	//nums := []int{1, 2, 9}
	nums := []int{2, 8, 7, 9}
	k := 2533
	n := getLargerButLessThanK(nums, k)
	fmt.Println(n)
}
