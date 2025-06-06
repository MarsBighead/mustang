package gstack

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	fmt.Println(ans)
}
