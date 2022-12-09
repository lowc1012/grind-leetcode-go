/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start


// Use a HashMap + doubly-linked list (container/list) to implement cache

// type Element struct {

// 		// The value stored with this element.
// 		Value any
// 		// contains filtered or unexported fields
// }

type Pair struct {
	Key int
	Val int
}

type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list.New(),
		make(map[int]*list.Element, capacity),
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exsit := this.m[key]; exsit {
		// move the node to front
		this.l.MoveToFront(node)
		return node.Value.(*list.Element).Value.(Pair).Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if cache hit
	if node, exsit := this.m[key]; exsit {
		// update the value of this node
		p := node.Value.(*list.Element).Value = Pair{
			key,
			value,
		}
		// move the node to front
		this.l.MoveToFront(node)

		// if cache miss
	} else {

		// push the new node into list, PushFront returns *Element
		ptr := this.l.PushFront(&list.Element{
			Value: Pair{
				Key: key,
				Val: value,
			},
		})

		// add the new item into map
		this.m[key] = ptr

		// if cache is overflow
		if this.l.Len() > this.capacity {
			// find the last node we want to delete
			lastKey := this.l.Back().Value.(*list.Element).Value.(Pair).Key
			// delete the item in map
			delete(this.m, lastNode.Key)

			// delete the node
			this.l.Remove(this.l.Back())
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

