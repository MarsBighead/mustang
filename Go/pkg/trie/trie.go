package trie

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd

}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func Run(actions []string, values [][]string) {
	if len(actions) != len(values) {
		return
	}
	var trie Trie
	for i, action := range actions {
		if i == 0 {
			if action != "Trie" {
				return
			} else {
				if len(values[i]) > 0 {
					return
				} else {
					trie = Constructor()
					fmt.Printf("null ")
				}
			}
			continue
		}
		switch action {
		case "insert":
			word := values[i][0]
			trie.Insert(word)
			fmt.Printf("null ")
		case "search":
			word := values[i][0]
			ok := trie.Search(word)
			fmt.Printf("%t ", ok)
		case "startsWith":
			word := values[i][0]
			ok := trie.StartsWith(word)
			fmt.Printf("%t ", ok)
		}
	}
	fmt.Println()
}
