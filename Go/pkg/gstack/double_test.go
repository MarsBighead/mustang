package gstack

import (
	"fmt"
	"strings"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	matrixes := [][][]byte{
		{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
		{{'0'}},
	}
	for _, matrix := range matrixes {
		n := maximalRectangle(matrix)
		fmt.Println(n)
	}
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
func TestSimplifyPath(t *testing.T) {
	s := "///"
	fmt.Printf("path: %#v\n", strings.Split(s, "/"))
	path := simplifyPath(s)
	fmt.Println(path)
}
func TestLongestValidParentheses(t *testing.T) {
	ss := []string{
		"()(())",
		"(()())",
		"(()",
		"",
		"()",
		")()())",
		"()(()",
		")()())()()(",
	}
	for _, s := range ss {
		n := longestValidParenthesesv1(s)
		fmt.Println(n)
	}
}
