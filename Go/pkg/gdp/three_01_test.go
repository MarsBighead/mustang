package gdp

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	numbers := [][]int{{1, 2, 3, 1}, {2, 7, 9, 3, 1}}
	for _, nums := range numbers {
		r := rob(nums)
		fmt.Println("amount=", r)
	}
}
