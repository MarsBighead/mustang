package bsearch

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	nums := []int{1, 2}
	head := array2List(nums)
	PrintList(head)
	head = removeNthFromEnd(head, 1)
	PrintList(head)

}
