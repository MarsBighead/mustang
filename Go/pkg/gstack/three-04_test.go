package gstack

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	ss := []string{
		"tree",
		"cccaaa",
		"Aabb",
		"loveleetcode",
	}
	for _, s := range ss {
		ans := frequencySort(s)
		fmt.Printf("  s=%s\n", s)
		fmt.Printf("ans=%s\n\n", ans)
	}
}
