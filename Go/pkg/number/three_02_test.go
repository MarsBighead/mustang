package number

import (
	"fmt"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	inputs := [][2]int{
		{5, 7},
		{0, 0},
		{1, 20},
	}
	for _, pos := range inputs {
		ans := rangeBitwiseAnd(pos[0], pos[1])
		fmt.Println("ans=", ans)
	}

}
