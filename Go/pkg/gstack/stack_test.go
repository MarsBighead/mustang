package gstack

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	s = removeDuplicates(s)
	fmt.Println(s)
}

func TestMultiply(t *testing.T) {
	num1 := "123"
	num2 := "456"
	ans := multiply(num1, num2)
	fmt.Println(ans)
}
