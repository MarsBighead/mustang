package bsearch

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}
	l1 := array2List(nums1)
	l2 := array2List(nums2)
	PrintList(l1)
	PrintList(l2)
	head := mergeTwoLists(l1, l2)
	PrintList(head)

}
