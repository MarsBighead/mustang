package bsearch

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var (
		hash    = make(map[*Node]*Node)
		current = head
		prev    = new(Node)
	)
	result := prev
	if head.Next == nil {
		node := &Node{Val: head.Val}
		if head.Random != nil {
			node.Random = node
		}
		return node
	}

	for current != nil {
		node := new(Node)
		node.Val = current.Val
		hash[current] = node
		prev.Next = node
		prev = prev.Next
		current = current.Next
	}
	current = head
	for current != nil {
		if current.Random != nil {
			node := hash[current]
			random := current.Random
			node.Random = hash[random]
		}
		current = current.Next
	}

	return result.Next

}

func Array2RandomList(array [][2]*int) *Node {
	if len(array) == 0 {
		return nil
	}
	var (
		head = &Node{Val: *array[0][0]}
		hash = make(map[int]*Node)
	)
	tail := head
	hash[0] = tail
	for i := 1; i < len(array); i++ {
		tail.Next = &Node{Val: *array[i][0]}
		hash[i] = tail.Next
		tail = tail.Next
	}
	for i := 0; i < len(array); i++ {
		if array[i][1] != nil {
			index := array[i][1]
			hash[i].Random = hash[*index]
		}
	}
	return head
}

func PrintRandomList(head *Node) {
	if head == nil {
		return
	}
	var (
		hash    = make(map[*Node]int)
		i       int
		current = head
	)
	for current != nil {
		hash[current] = i
		i++
		current = current.Next
	}
	current = head
	for current != nil {
		if current.Random != nil {
			y := hash[current.Random]
			fmt.Printf("[%d %d] ", current.Val, y)
		} else {
			fmt.Printf("[%d null] ", current.Val)
		}
		current = current.Next
	}
	fmt.Println()

}
