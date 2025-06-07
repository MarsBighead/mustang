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

func TestRotateRight(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := ArrayToList(nums)
	PrintList(head)
	head = rotateRight(head, 2)
	PrintList(head)

}
func TestDeleteAllDuplicates(t *testing.T) {

	//nums := []int{1, 1, 2, 3, 4, 5}
	nums := []int{1, 2, 3, 3, 4, 4, 5}
	head := ArrayToList(nums)
	fmt.Printf("before: %d\n", head.Val)
	PrintList(head)
	head = deleteAllDuplicates(head)
	fmt.Printf("after: %d\n", head.Val)
	PrintList(head)
}
func TestDeleteDuplicates(t *testing.T) {

	nums := []int{1, 1, 2, 3, 4, 5}
	nums = []int{1, 1, 1}
	head := ArrayToList(nums)
	PrintList(head)
	deleteDuplicates(head)
	PrintList(head)
}

func TestPartition(t *testing.T) {
	numbers := []int{1, 4, 3, 2, 5, 2}
	x := 3
	head := ArrayToList(numbers)
	ans := partition(head, x)
	PrintList(ans)
}

func TestPartitionv1(t *testing.T) {
	numbers := []int{1, 4, 3, 2, 5, 2}
	x := 3
	head := ArrayToList(numbers)
	ans := partitionv1(head, x)
	PrintList(ans)
}
func TestReverseBetween(t *testing.T) {
	nums := []int{3, 5}
	head := ArrayToList(nums)
	PrintList(head)
	head = reverseBetween(head, 1, 2)
	PrintList(head)
}
