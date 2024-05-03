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

func TestIsValid(t *testing.T) {
	s := "{[]}"
	//s = "()))"
	//s = "[[[]"
	//s = "()"
	//s = "()[]{}"
	ok := isValid(s)
	fmt.Println(ok)
}
