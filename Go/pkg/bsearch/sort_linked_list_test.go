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

func TestSortList(t *testing.T) {
	nums := []int{4, 2, 1, 3}
	head := array2List(nums)
	PrintList(head)
	sorted := sortList(head)
	PrintList(sorted)
}

func TestMergeKLists(t *testing.T) {
	//nums := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	nums := [][]int{{}, {1, 4, 5}, {}, {2, 6}}
	var lists []*ListNode
	for _, row := range nums {
		head := array2List(row)
		lists = append(lists, head)

	}
	sorted := mergeKListsV2(lists)
	PrintList(sorted)
}

func TestMergeSort(t *testing.T) {
	nums := []int{4, 2, 1, 3, 5}
	head := array2List(nums)
	PrintList(head)
	sorted := mergeSort(head)
	fmt.Println("result:")
	PrintList(sorted)
}
