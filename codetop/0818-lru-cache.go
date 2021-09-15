package codetop

type DLinkedNode struct {
	key  int
	val  int
	prev *DLinkedNode
	next *DLinkedNode
}

type LRUCache struct {
	cache    map[int]*DLinkedNode
	recode   *DLinkedNode
	capacity int
	size     int
}

func (head *DLinkedNode) moveNodeToHead(node *DLinkedNode) {
	if node == head.next {
		return
	}
	head.removeNode(node)
	head.addToHead(node)
}

func (head *DLinkedNode) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (head *DLinkedNode) addToHead(node *DLinkedNode) {
	node.next = head.next
	node.prev = head
	head.next = node
	node.next.prev = node
}

func ConstructorLRU(capacity int) LRUCache {
	head := &DLinkedNode{
		val:  -1,
		prev: nil,
		next: nil,
	}
	head.next = head
	head.prev = head
	return LRUCache{
		cache:    make(map[int]*DLinkedNode),
		recode:   head,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.recode.moveNodeToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.recode.moveNodeToHead(node)
	} else {
		if this.size == this.capacity {
			node := this.recode.prev
			delete(this.cache, node.key)
			node.key = key
			node.val = value
			this.recode.moveNodeToHead(node)
			this.cache[key] = node
		} else {
			node := &DLinkedNode{key: key, val: value, prev: nil, next: nil}
			this.recode.addToHead(node)
			this.cache[key] = node
			this.size++
		}
	}
}
