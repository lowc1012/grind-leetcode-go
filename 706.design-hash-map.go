/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
type Node struct {
	key, val int
	next     *Node
}

type MyHashMap struct {
	buckets []*Node
}

func Constructor() MyHashMap {
	return MyHashMap{make([]*Node, 1000)}
}

func (m *MyHashMap) hash(key int) int {
	return key % len(m.buckets)
}

func (m *MyHashMap) Put(key, value int) {
	idx := m.hash(key)
    // new node in the bucket
	if m.buckets[idx] == nil {
		m.buckets[idx] = &Node{key, value, nil}
		return
	}

	for node := m.buckets[idx]; ; node = node.next {
		if node.key == key {
			node.val = value
			return
		}
		if node.next == nil {
			node.next = &Node{key, value, nil}
			return
		}
	}
}

func (m *MyHashMap) Get(key int) int {
	for node := m.buckets[m.hash(key)]; node != nil; node = node.next {
		if node.key == key {
			return node.val
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	idx := m.hash(key)
	if m.buckets[idx] == nil {
		return
	}

	if m.buckets[idx].key == key {
		m.buckets[idx] = m.buckets[idx].next
		return
	}

	for node := m.buckets[idx]; node.next != nil; node = node.next {
		if node.next.key == key {
			node.next = node.next.next
			return
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

