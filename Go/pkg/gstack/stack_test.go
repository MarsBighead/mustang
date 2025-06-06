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

func TestCalculte(t *testing.T) {
	s := " 0-2147483647"
	r := calculate(s)
	fmt.Println(s, r)
}

func TestCalcultev2(t *testing.T) {
	s := " 0-2147483647"
	r := calculatev2(s)
	fmt.Println(s, r)
}

func TestMultiply(t *testing.T) {
	num1 := "123"
	num2 := "456"
	ans := multiply(num1, num2)
	fmt.Println(ans)
}
