package hash

import (
	"fmt"
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	cpdomainss := [][]string{
		{"9001 discuss.leetcode.com"},
		{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
	}
	for _, cpdomains := range cpdomainss {
		ans := subdomainVisits(cpdomains)
		fmt.Printf("ans=%#v\n\n", ans)
	}
}

func TestIsHappy(t *testing.T) {
	nums := []int{19, 2}
	for _, n := range nums {
		ok := isHappy(n)
		fmt.Printf("%d:%t\n\n", n, ok)
	}

}
