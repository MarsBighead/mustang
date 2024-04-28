package linked

import (
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
