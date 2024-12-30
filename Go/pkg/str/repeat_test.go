package str

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	ss := []string{
		"abab",
		"aba",
		"abcabcabcabc",
	}
	for _, s := range ss {
		fmt.Printf("org=%s\n", s)
		ok := repeatedSubstringPattern(s)
		fmt.Printf(" ok=%t\n\n", ok)
	}

}
