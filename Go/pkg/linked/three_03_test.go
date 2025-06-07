package linked

import "testing"

func TestOddEvenList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := ArrayToList(nums)
	PrintList(head)

	head = oddEvenList2(head)
	PrintList(head)
}
