package gstack

import (
	"fmt"
	"testing"
)

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
