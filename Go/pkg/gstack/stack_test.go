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
