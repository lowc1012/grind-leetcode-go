/*
 * @lc app=leetcode id=705 lang=golang
 *
 * [705] Design HashSet
 */

// @lc code=start

type MyHashSet struct {
	items map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	if _, exist := this.items[key]; !exist {
		this.items[key] = struct{}{}
	}
}

func (this *MyHashSet) Remove(key int) {
	if _, exist := this.items[key]; exist {
		delete(this.items, key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	_, exist := this.items[key]
	return exist
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

