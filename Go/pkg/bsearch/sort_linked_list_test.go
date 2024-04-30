package bsearch

import (
	"fmt"
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

func TestHasCycle(t *testing.T) {
	nums := []int{1, 2}
	head := array2List(nums)
	PrintList(head)
	ok := hasCycle(head)
	fmt.Println(ok)
}
