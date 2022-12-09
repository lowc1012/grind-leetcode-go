/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
type Node struct {
	key  int
	val  int
	next *Node
}

type MyHashMap struct {
	items []*Node
	size  int
}

func Constructor() MyHashMap {
	size := 1000
	return MyHashMap{
		make([]*Node, size),
		size,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % this.size
	if this.items[index] == nil {
		this.items[index] = &Node{
			key: key,
			val: value,
		}
		return
	}

	item := this.items[index]
	for item != nil {
		if item.key == key {
			item.val = value
			return
		}

		if item.next == nil {
			item.next = &Node{key: key, val: value}
			return
		}

		item = item.next
	}
}

func (this *MyHashMap) Get(key int) int {
	index := key % this.size
	if this.items[index] == nil {
		return -1
	}

	item := this.items[index]

	for item != nil {
		if item.key == key {
			return item.val
		}

		item = item.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % this.size
	if this.items[index] == nil {
		return
	} else {

		if this.items[index].key == key {
			this.items[index] = this.items[index].next
			return
		}

		item := this.items[index]
		for item.next != nil {
			if item.next.key == key {
				item.next = item.next.next
				return
			}
			item = item.next
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

