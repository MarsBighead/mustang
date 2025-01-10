package opt

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	sstrs := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}
	for _, strs := range sstrs {
		ans := groupAnagrams(strs)
		fmt.Printf("ans=%#v\n\n", ans)
	}
}

func TestReverseWords(t *testing.T) {
	ss := []string{
		"  hello world  ",
		"the sky is blue",
		"  asdasd df f",
		"EPY2giL",
		" EPY2giL",
		"EPY2giL ",
	}
	for _, s := range ss {
		fmt.Printf("org=|%s|\n", s)
		ans := reverseWordsv1(s)
		fmt.Printf("ans=|%s|\n###\n\n", ans)
	}
}

func TestReverseWords3(t *testing.T) {
	ss := []string{
		"Let's take LeetCode contest",
		"I love ux",
	}
	for _, s := range ss {
		//fmt.Println(len(s), s)
		ans := reverseWords3(s)
		fmt.Println(len(ans), ans)
	}
}
