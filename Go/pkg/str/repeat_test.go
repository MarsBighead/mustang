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

func TestMaxRepeating(t *testing.T) {
	ss := []string{
		"ababc",
		"ababc",
		"ababc",
		"aaabaaaabaaabaaaabaaaabaaaabaaaaba",
		"aaa",
		"baba",
	}
	words := []string{
		"ab",
		"ba",
		"ac",
		"aaaba",
		"a",
		"b",
	}
	for i, s := range ss {
		fmt.Printf("org=%s\n", s)
		num := maxRepeatingv1(s, words[i])
		fmt.Printf("num=%d\n\n", num)
	}

}

func TestStringMatching(t *testing.T) {
	wordss := [][]string{
		{"mass", "as", "hero", "superhero"},
		{"leetcode", "et", "code"},
	}
	for _, words := range wordss {
		ans := stringMatchingv1(words)
		fmt.Println(ans)
	}
}
