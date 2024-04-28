package linked

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	ln := ArrayToList(head)
	PrintList(ln)
	/*if ln != nil {
		fmt.Printf("x：%#v\n", ln.Val)
	}
	for ln.Next != nil {
		ln = ln.Next
		fmt.Println("List value=", ln.Val)
	}*/

}

func TestDeleteDuplicates(t *testing.T) {

	nums := []int{1, 1, 2, 3, 4, 5}
	nums = []int{1, 1, 1}
	head := ArrayToList(nums)
	PrintList(head)
	deleteDuplicates(head)
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
