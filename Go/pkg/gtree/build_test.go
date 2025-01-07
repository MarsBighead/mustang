package gtree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTreev1(preorder, inorder)

	fmt.Printf("%#v\n", root)
	fmt.Printf("%#v\n", preorder)
	order := preorderTraversal(root)
	fmt.Printf("%#v\n", order)
}

func TestBuildTreev1(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 6, 7, 3, 8, 9}
	inorder := []int{4, 2, 6, 5, 7, 1, 9, 8, 3}
	root := buildTreev1(preorder, inorder)

	fmt.Printf("%#v\n", root)
	fmt.Printf("%#v\n", preorder)
	order := preorderTraversal(root)
	fmt.Printf("%#v\n", order)
}

func TestBuildTree2(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree2(inorder, postorder)

	fmt.Printf("%#v\n", inorder)
	fmt.Printf("%#v\n", postorder)
	order := preorderTraversal(root)
	fmt.Printf("%#v\n", order)
}
