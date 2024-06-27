package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int                   // 缓存容量
	cache    map[int]*list.Element // 哈希表
	lruList  *list.List            // 双向链表
}

type entry struct {
	key   int
	value int
}

// NewLRUCache 初始化一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lruList:  list.New(),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if elem, found := c.cache[key]; found {
		c.lruList.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return -1, false
}

func (c *LRUCache) Put(key int, value int) {
	if elem, found := c.cache[key]; found {
		c.lruList.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.lruList.Len() == c.capacity {
		elem := c.lruList.Back()
		if elem != nil {
			c.lruList.Remove(elem)
			delete(c.cache, elem.Value.(*entry).key)
		}
	}

	newElem := c.lruList.PushFront(&entry{key, value})
	c.cache[key] = newElem
}

func main() {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回 1

	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)

	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回 3
	fmt.Println(cache.Get(4)) // 返回 4
}
