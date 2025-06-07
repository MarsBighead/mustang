package linked

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := ArrayToList(nums)
	head = reverseKGroup(head, 2)
	fmt.Printf("result: ")
	PrintList(head)
	nums = []int{1}
	head = ArrayToList(nums)
	head = reverseKGroup(head, 1)
	fmt.Printf("result: ")
	PrintList(head)

}
