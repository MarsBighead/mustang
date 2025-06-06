package number

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	numbers := [][]int{{1, 1, 0, 1, 1, 1}, {1, 0, 1, 1, 0, 1}}
	for _, nums := range numbers {
		ans := findMaxConsecutiveOnes(nums)
		fmt.Println("ans=", ans)
	}
}
