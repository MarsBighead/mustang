package graph

import (
	"fmt"
	"testing"
)

func TestFindOrder(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	ans := findOrder(numCourses, prerequisites)
	fmt.Printf("%#v\n", ans)
}
