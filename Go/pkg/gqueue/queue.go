package gqueue

import (
	"fmt"
	"sort"
)

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

// https://leetcode.cn/problems/kth-largest-element-in-a-stream/description/
// 703. 数据流中的第 K 大元素
type KthLargest struct {
	k    int
	nums []int
}

func KConstructor(k int, nums []int) KthLargest {
	length := len(nums)
	sort.Ints(nums)
	if length >= k {
		nums = nums[length-k:]
	}

	return KthLargest{
		k,
		nums,
	}

}

// add insert sort
func (this *KthLargest) Add(val int) int {
	length := len(this.nums)
	if length == this.k {
		if val > this.nums[0] {
			this.nums[0] = val
			for i := 1; i < length; i++ {
				if this.nums[i] < this.nums[i-1] {
					this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
				} else {
					break
				}
			}
		}

	} else if length < this.k {
		this.nums = append(this.nums, val)
		length = len(this.nums)
		for i := length - 2; i >= 0; i-- {
			if this.nums[i+1] < this.nums[i] {
				this.nums[i+1], this.nums[i] = this.nums[i], this.nums[i+1]
			} else {
				break
			}
		}
	}
	//sort.Ints(this.nums)
	fmt.Println(this.nums)

	if len(this.nums) > this.k {
		this.nums = this.nums[len(this.nums)-this.k:]
		return this.nums[0]
	} else if len(this.nums) == this.k {
		return this.nums[0]
	} else {
		return 0
	}

}
