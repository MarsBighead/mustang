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
