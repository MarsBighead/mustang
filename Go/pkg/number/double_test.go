package number

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	numbers := [][]int{{1, 3}, {2, 6}, {8, 1}, {15, 18}}
	ans := merge(numbers)
	fmt.Printf("%#v\n", ans)
	numbers = [][]int{{1, 4}, {2, 3}}
	ans = merge(numbers)
	fmt.Printf("%#v\n", ans)
}
