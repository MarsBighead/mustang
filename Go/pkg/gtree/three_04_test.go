package gtree

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 6},
		},
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
	}
	ans := deleteNode(root, 5)
	order := preorderTraversal(ans)
	fmt.Printf("%#v\n", order)
}
