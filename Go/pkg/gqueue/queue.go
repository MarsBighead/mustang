package gqueue

import "fmt"

// 定义一个简单的循环队列
type MyCircularQueue struct {
	capacity int
	elements []int
	head     int
	tail     int
}

// 初始化循环队列
func NewCircularQueue(capacity int) *MyCircularQueue {
	return &MyCircularQueue{
		capacity: capacity + 1,
		elements: make([]int, capacity),
		head:     0,
		tail:     0,
	}
}

// 入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

// 出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	//value := this.elements[this.head]
	for i := 0; i < this.capacity-2; i++ {
		this.elements[i] = this.elements[i+1]
	}
	this.elements[this.capacity-2] = 0
	this.tail = (this.tail - 1) % this.capacity
	fmt.Println("tail=", this.tail)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	fmt.Println(this.tail, this.elements)
	return this.elements[this.tail-1]
}

// 队列是否已满
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.capacity == this.head
}

// 队列是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}
