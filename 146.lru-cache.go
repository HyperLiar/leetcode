/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	cap  int
	m    map[int]*JNode
	head *JNode
	tail *JNode
}

type JNode struct {
	key  int
	val  int
	prev *JNode
	next *JNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, m: make(map[int]*JNode)}
}

func (this *LRUCache) addToList(n *JNode) {
	n.prev = nil
	n.next = this.head

	if this.head != nil {
		this.head.prev = n
	}

	this.head = n
	if this.tail == nil {
		this.tail = n
	}
}

func (this *LRUCache) removeFromList(n *JNode) {
	if n != this.head {
		n.prev.next = n.next
	} else {
		this.head = n.next
	}
	if n != this.tail {
		n.next.prev = n.prev
	} else {
		this.tail = n.prev
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.removeFromList(n)
		this.addToList(n)
		return n.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.val = value
		this.removeFromList(node)
		this.addToList(node)
		return
	} else {
		node = &JNode{key: key, val: value}
		this.m[key] = node
		this.addToList(node)
	}
	if len(this.m) > this.cap {
		delete(this.m, this.tail.key)
		this.removeFromList(this.tail)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

