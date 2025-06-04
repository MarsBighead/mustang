package base

import "math/rand"

// https://leetcode.cn/problems/insert-delete-getrandom-o1/description
// 380. O(1) 时间插入、删除和获取随机元素
type RandomizedSet struct {
	elements map[int]int // 存储元素值和其在切片中的索引
	items    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		elements: make(map[int]int),
		items:    make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.elements[val]; ok {
		return false
	}
	this.elements[val] = len(this.items)
	this.items = append(this.items, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.elements[val]; ok {
		// 从map中移除元素索引
		delete(this.elements, val)
		lastIndex := len(this.items) - 1
		if lastIndex != index {
			this.items[index] = this.items[lastIndex]
			this.elements[this.items[index]] = index
		}
		this.items = this.items[:lastIndex]
		return true
	}
	// 元素不存在，返回false
	return false

}

func (this *RandomizedSet) GetRandom() int {
	if len(this.items) == 0 {
		return 0
	}
	randomIndex := rand.Intn(len(this.items))
	return this.items[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
