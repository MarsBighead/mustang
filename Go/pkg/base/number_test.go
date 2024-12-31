package base

import (
	"fmt"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	numbers := []int{
		100,
		-7,
		-8,
	}
	for _, num := range numbers {
		fmt.Printf("num=%d\n", num)
		ans := convertToBase7(num)
		fmt.Printf("base7=%s\n\n", ans)
	}
}
