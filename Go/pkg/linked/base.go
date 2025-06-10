package linked

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
交叉链表
*/
// https://leetcode.cn/problems/3u1WK4/description/
// LCR 023. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	// 链表拼接，最后一段肯定是相同的
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}

	}
	return a
}

func ArrayToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head

}

func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println()
		return
	}
	fmt.Print(head.Val, " ")
	for head.Next != nil {
		head = head.Next
		fmt.Print(head.Val, " ")
	}
	fmt.Println()
}

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
