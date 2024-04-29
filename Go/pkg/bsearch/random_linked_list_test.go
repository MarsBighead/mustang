package bsearch

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	head := &Node{
		Val: 3,
	}
	head.Next = &Node{Val: 3, Random: head}
	head.Next.Next = &Node{Val: 3}

	fmt.Printf("Raw   : ")
	PrintRandomList(head)
	result := copyRandomList(head)
	fmt.Printf("Copied: ")
	PrintRandomList(result)

	nums := [][]int{{7}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	array := make([][2]*int, len(nums))
	for i, row := range nums {
		arr := [2]*int{&row[0], nil}
		if len(row) == 2 {
			arr[1] = &row[1]
		}
		array[i] = arr
	}
	fmt.Printf("%#v\n", array)
	head = Array2RandomList(array)
	PrintRandomList(head)

}
