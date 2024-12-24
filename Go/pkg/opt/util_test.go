package opt

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	s := 128
	s = 18
	r := IsPowerOfTwo(s)
	fmt.Println(s, r)
}

func TestCtz64(t *testing.T) {
	var s uint64 = 7

	r := Ctz64(s)
	fmt.Println(s, r)
}
