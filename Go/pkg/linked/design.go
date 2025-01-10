package linked

import "fmt"

type Node struct {
	value int
	next  *Node
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
type MyLinkedList struct {
	len  int
	head *Node
}

func Constructor() MyLinkedList {
	this := new(MyLinkedList)
	this.len = 0
	this.head = nil
	return *this

}

func (this *MyLinkedList) Get(index int) int {
	if index > this.len-1 {
		return -1
	}
	cur := this.head
	for i := 0; i != index; i++ {
		cur = cur.next
	}
	return cur.value

}

func (this *MyLinkedList) AddAtHead(val int) {
	node := new(Node)
	node.value = val

	if this.head == nil {
		this.head = node
	} else {
		node.next = this.head
		this.head = node
	}
	this.len++

}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.head
	for i := 0; i < this.len-1; i++ {
		cur = cur.next
	}
	node := new(Node)
	node.value = val
	if cur == nil {
		this.head = node
	} else {
		cur.next = node
	}

	this.len++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}

	if index == 0 {

		this.AddAtHead(val)
		return
	}
	if index == this.len {

		this.AddAtTail(val)
		return
	}

	cur := this.head

	for i := 0; i != index-1; i++ {
		cur = cur.next
	}
	node := new(Node)
	node.value = val
	node.next = cur.next
	cur.next = node
	this.len++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head
	if index > this.len-1 {
		return
	}

	if index == 0 {
		this.head = this.head.next
	}

	fmt.Println(this.len)

	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	if index == this.len-1 {

		cur.next = nil
		this.len--
		return
	}

	cur.next = cur.next.next
	this.len--
}
