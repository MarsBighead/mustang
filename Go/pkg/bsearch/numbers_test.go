package bsearch

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	ss := []string{
		"42",
		"   -042",
		"+-12",
		/*"1337c0d3",
		"0-1",
		"words and 987",
		"-91283472332",*/
	}
	for _, s := range ss {
		n := myAtoi(s)
		fmt.Println(n)
	}
}

func TestNumDecodings(t *testing.T) {
	ss := []string{
		"12",
		"226",
		"06",
		"42",
		"250",
	}
	for _, s := range ss {
		n := numDecodings(s)
		fmt.Println(n)
	}
}
