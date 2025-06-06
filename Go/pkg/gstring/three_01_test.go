package gstring

import (
	"fmt"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	ss := []string{
		"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		"AAAAAAAAAAAAA",
		"AAAAAAAAAAA",
	}
	for _, s := range ss {
		ans := findRepeatedDnaSequences(s)
		fmt.Println(ans)
	}
}
