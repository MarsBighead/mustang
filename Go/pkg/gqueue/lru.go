package gqueue

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[interface{}]*list.Element
	list     *list.List
}
type cacheItem struct {
	key   interface{}
	value interface{}
}

func LruConstructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if element, exists := this.cache[key]; exists {
		this.list.MoveToFront(element) // 将访问的元素移到链表头部，表示最近使用
		return element.Value.(*cacheItem).value.(int)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, exists := this.cache[key]; exists {
		// 如果键已存在，更新其值并移动到链表头部表示最近使用
		this.list.MoveToFront(element)
		element.Value.(*cacheItem).value = value
		return
	}

	// 如果键不存在，创建一个新的元素并添加到链表头部和缓存中
	element := this.list.PushFront(&cacheItem{key: key, value: value})
	this.cache[key] = element

	// 如果缓存满了，移除链表尾部的元素（即最久未使用的）和对应的缓存项
	if this.list.Len() > this.capacity {
		lastElement := this.list.Back()                        // 获取链表尾部元素（最久未使用的）
		this.list.Remove(lastElement)                          // 从链表中移除该元素
		delete(this.cache, lastElement.Value.(*cacheItem).key) // 从缓存中删除对应的键值对
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
