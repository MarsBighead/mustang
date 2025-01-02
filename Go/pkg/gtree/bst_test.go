package gtree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	trees := []*TreeNode{
		{Val: 1},
		{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
	}
	for _, root := range trees {
		ok := isValidBST(root)
		fmt.Printf("isValidBST: %t\n", ok)
	}
}
