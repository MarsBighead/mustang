package gqueue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	this := NewCircularQueue(3)
	ok := this.EnQueue(1)
	ok = this.EnQueue(2)
	ok = this.EnQueue(3)
	fmt.Println(ok)
	fmt.Printf("head=%d, tail=%d;queue: %v\n", this.head, this.tail, this.elements)
	ok = this.EnQueue(4)
	ok = this.DeQueue()
	ok = this.DeQueue()
	fmt.Println(ok)
	fmt.Printf("head=%d, tail=%d;queue: %v\n", this.head, this.tail, this.elements)
	ok = this.EnQueue(4)
	fmt.Printf("head=%d, tail=%d;queue: %v\n", this.head, this.tail, this.elements)
	val := this.Rear()
	fmt.Println("value=", val)
}

func TestKConstructor(t *testing.T) {
	nums := []int{5, -1}
	k := 3
	kl := KConstructor(k, nums)
	result := []int{}
	result = append(result, kl.Add(2))
	result = append(result, kl.Add(1))
	result = append(result, kl.Add(3))
	result = append(result, kl.Add(4))
	fmt.Println("result=", result)
}
