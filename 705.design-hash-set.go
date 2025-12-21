/*
 * @lc app=leetcode id=705 lang=golang
 *
 * [705] Design HashSet
 */

// @lc code=start

type Node struct {
    key  int
    next *Node
}

type MyHashSet struct {
    set *Node
}

func Constructor() MyHashSet {
    return MyHashSet{
        set: nil,
    }
}

func (this *MyHashSet) Add(key int) {
    if this.set == nil {
        this.set = &Node{
            key: key,
            next: nil,
        }
        return
    }
    for node := this.set; ; node = node.next {
        if node.key == key {
            break
        }
        if node.next == nil {
            node.next = &Node{
                key: key,
                next: nil,
            }
            return
        }
    }
}

func (this *MyHashSet) Remove(key int) {
    if this.set == nil {
        return
    }
    if node := this.set; node.key == key {
        this.set = node.next
        return
    }
    for node := this.set; node.next != nil; node = node.next {
        if node.next.key == key {
            node.next = node.next.next
            return
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    for node := this.set; node != nil ; node = node.next {
        if node.key == key {
            return true
        }
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

