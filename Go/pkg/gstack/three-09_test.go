package gstack

import (
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	hash := make(map[string][][]int)
	hash = map[string][][]int{
		"3": {{2, 1, 0}, {1, 2, 0}},
		"1": {{1, 2, 3, 4, 5}, {4, 5, 3, 2, 1}},
		"2": {{1, 2, 3, 4, 5}, {4, 3, 5, 1, 2}},
	}
	for key, val := range hash {
		fmt.Println(key, val[0], val[1])
		pushed, popped := val[0], val[1]
		ok := validateStackSequences(pushed, popped)
		fmt.Println("### ", ok)
	}
}

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
