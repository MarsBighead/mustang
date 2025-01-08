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

func TestConstructFromPrePost(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(preorder, postorder)

	fmt.Printf("%#v\n", preorder)
	fmt.Printf("%#v\n", postorder)
	preorder = []int{1, 2}
	postorder = []int{2, 1}
	root = constructFromPrePost(preorder, postorder)
	fmt.Printf("     %#v\nleft %#v\nright %#v\n", root, root.Left, root.Right)
	order := preorderTraversal(root)
	fmt.Printf("%#v\n", order)
}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	ans := inorderTraversal(root)
	fmt.Printf("ans=%#v\n", ans)
}
