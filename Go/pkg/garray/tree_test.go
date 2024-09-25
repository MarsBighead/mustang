package garray

import (
	"fmt"
	"testing"
)

func TestConstructTree(t *testing.T) {
	nums := []int{5, 2}
	mask := 1 << 61
	fmt.Println(mask)
	fmt.Println(mask >> 60)

	tree := Construct(nums)
	fmt.Printf("%#v\n", tree)
	printTree(tree)
}

func TestPostorderTraversal(t *testing.T) {
	nums := []int{3, 1, 2}
	//mask := 1 << 61
	//fmt.Println(mask)
	//fmt.Println(mask >> 60)

	tree := Construct(nums)
	result := postorderTraversalv1(tree)
	fmt.Printf("ninhao: %#v\n", result)

}

func TestParseBrConifgOptions(t *testing.T) {
	body := "[--check-requirements=false]"
	options, err := parseBrConifgOptions(body)
	fmt.Println(options, err)
}
