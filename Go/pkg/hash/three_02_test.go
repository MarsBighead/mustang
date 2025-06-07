package hash

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	nums := []int{19, 2}
	for _, n := range nums {
		ok := isHappy(n)
		fmt.Printf("%d:%t\n\n", n, ok)
	}

}
