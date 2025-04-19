package lld

// Node for doubly linked list
type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// remove a node from linked list
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// add a node right after head (most recent position)
func (this *LRUCache) addToFront(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.remove(node)
		this.addToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.remove(node)
		this.addToFront(node)
	} else {
		if len(this.cache) == this.capacity {
			// remove LRU node from tail
			lru := this.tail.prev
			this.remove(lru)
			delete(this.cache, lru.key)
		}
		newNode := &Node{key: key, value: value}
		this.addToFront(newNode)
		this.cache[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
