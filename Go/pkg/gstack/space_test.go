package gstack

import (
	"fmt"
	"testing"
)

func TestKClosest(t *testing.T) {
	ks := []int{1}
	apoints := [][][]int{
		{{1, 3}, {-2, 2}},
	}
	for i, k := range ks {
		points := apoints[i]
		ans := kClosestv1(points, k)
		fmt.Println("result:")
		for _, point := range ans {
			fmt.Printf("%v ", point)
		}
		fmt.Println()
	}
}
