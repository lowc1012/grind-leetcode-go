/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

import "container/list"

// @lc code=start

// Use a HashMap + doubly-linked list (container/list) to implement cache

// type Element struct {

// 		// The value stored with this element.
// 		Value any
// 		// contains filtered or unexported fields
// }

type Pair struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	lruList  *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lruList:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, exist := this.cache[key]; exist {
		this.lruList.MoveToFront(ele)
		return ele.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, exist := this.cache[key]; exist {
		this.lruList.MoveToFront(ele)
		ele.Value = Pair{key, value}
	} else {
		if len(this.cache) >= this.capacity {
			theOldest := this.lruList.Back()
			delete(this.cache, theOldest.Value.(Pair).key)
			this.lruList.Remove(theOldest)
		}
		this.cache[key] = this.lruList.PushFront(Pair{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

