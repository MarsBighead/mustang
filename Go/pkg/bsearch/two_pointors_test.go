package bsearch

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	nums := []int{1, 2}
	head := array2List(nums)
	PrintList(head)
	//head = removeNthFromEnd(head, 1)
	head = removeNthFromEnd2(head, 2)
	PrintList(head)

}

func TestDetectCycle(t *testing.T) {
	//nums := []int{1, 2}
	//head := array2List(nums)
	//AddCycle2List(head, 0)
	nums := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	fmt.Println(len(nums))
	head := array2List(nums)
	AddCycle2List(head, 24)

	node := detectCycle2(head)
	fmt.Println("node.Val=", node.Val)
	//PrintList(head)
	//head = removeNthFromEnd(head, 1)
	//head = removeNthFromEnd2(head, 2)
	//PrintList(head)

}

func TestReorderList(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5}
	head := array2List(nums)
	PrintList(head)
	//head = removeNthFromEnd(head, 1)
	reorderList(head)
}

func TestAddTwoNumbersv2(t *testing.T) {

	nums1 := []int{9, 9, 9}
	nums2 := []int{1}
	l1 := array2List(nums1)
	l2 := array2List(nums2)
	l1 = addTwoNumbersv2(l1, l2)
	PrintList(l1)
	//head = removeNthFromEnd(head, 1)
}
