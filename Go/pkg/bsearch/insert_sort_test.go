package bsearch

import (
	"fmt"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	nums := []int{4, 2, 1, 3, 5}
	head := array2List(nums)
	PrintList(head)
	sorted := insertionSortList(head)
	fmt.Println("result:")
	PrintList(sorted)
}
