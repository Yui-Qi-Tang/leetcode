package sol

// https://leetcode.com/problems/lru-cache/

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

// LRUCache ...
type LRUCache struct {
	data map[int]*node
	head *node
	tail *node
	cap  int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{data: make(map[int]*node, capacity), cap: capacity}
}

// Get ...
func (lr *LRUCache) Get(key int) int {
	if n, exist := lr.data[key]; exist {
		lr.toTop(n)
		return n.value
	}
	return -1
}

// Put ...
func (lr *LRUCache) Put(key int, value int) {

	// insert if key does not exist, and move node to top
	// update if key does exist, and move node to top
	// remove if capacity is full

	if n, exist := lr.data[key]; exist {
		n.value = value
		lr.toTop(n)
		return
	}

	n := &node{key: key, value: value}
	lr.data[key] = n

	if lr.head == nil {
		lr.head = n
		lr.tail = n
		return
	}

	lr.add(n)

	if len(lr.data) > lr.cap {
		lr.deleteTail()
	}
}

func (lr *LRUCache) add(n *node) {
	lr.head.prev = n
	n.next = lr.head
	lr.head = n
}

func (lr *LRUCache) toTop(n *node) {

	// if head -> just update
	if n == lr.head {
		return
	}

	// if tail -> set tail
	if n == lr.tail {
		lr.tail = n.prev
	}

	// move n to top
	n.prev.next = n.next

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.prev = nil

	lr.head.prev = n
	n.next = lr.head

	lr.head = n
}

func (lr *LRUCache) deleteTail() {
	tail := lr.tail
	delete(lr.data, tail.key) // remove data from cache
	lr.tail = tail.prev
	lr.tail.next = nil
}
