package linked

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := ArrayToList(nums)
	PrintList(head)
	head = removeElements(head, 6)
	PrintList(head)
}
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
func TestIsPalindrome(t *testing.T) {
	nums := []int{1, 2, 2, 1}
	head := ArrayToList(nums)
	PrintList(head)
	ok := isPalindrome2(head)
	PrintList(head)
	fmt.Println("ok=", ok)
}
