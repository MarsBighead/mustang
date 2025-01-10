package trie

import (
	"testing"
)

func TestRun(t *testing.T) {
	actions := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	values := [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}
	Run(actions, values)
}
