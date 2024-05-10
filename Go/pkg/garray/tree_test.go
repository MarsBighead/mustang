package garray

import (
	"fmt"
	"testing"
)

func TestConstructTree(t *testing.T) {
	nums := []int{5, 2}
	tree := Construct(nums)
	fmt.Printf("%#v\n", tree)
	printTree(tree)
}
