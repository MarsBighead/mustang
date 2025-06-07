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
