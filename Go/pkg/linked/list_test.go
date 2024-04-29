package linked

import (
	"fmt"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := ArrayToList(nums)
	PrintList(head)
	head = oddEvenList2(head)
	PrintList(head)
}

func TestIsPalindrome(t *testing.T) {
	nums := []int{1, 2, 2, 1}
	head := ArrayToList(nums)
	PrintList(head)
	ok := isPalindrome2(head)
	PrintList(head)
	fmt.Println("ok=", ok)
}
