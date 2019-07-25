package lru

// 哈希表 + 双向链表

type DLinkedNode struct {
	Key  int
	Val  int
	Pre  *DLinkedNode
	Next *DLinkedNode
}

type LRUCache struct {
	data     map[int]*DLinkedNode
	capacity int
	count    int
	head     *DLinkedNode
	tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedNode{}
	tail := &DLinkedNode{}
	head.Next = tail
	tail.Pre = head

	data := make(map[int]*DLinkedNode)
	return LRUCache{
		data:     data,
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if !ok {
		newNode := &DLinkedNode{
			Key: key,
			Val: value,
		}
		this.data[key] = newNode
		this.addNode(newNode)
		this.count++
		if this.count > this.capacity {
			tail := this.popTail()
			delete(this.data, tail.Key)
			this.count--
		}
	} else {
		node.Val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	prev := node.Pre
	next := node.Next
	prev.Next = next
	next.Pre = prev

	hNext := this.head.Next
	this.head.Next = node
	node.Pre = this.head

	node.Next = hNext
	hNext.Pre = node
}

func (this *LRUCache) addNode(node *DLinkedNode) {
	next := this.head.Next
	this.head.Next = node
	node.Pre = this.head
	node.Next = next
	next.Pre = node
}

func (this *LRUCache) popTail() *DLinkedNode {
	node := this.tail.Pre
	pre := node.Pre
	pre.Next = this.tail
	this.tail.Pre = pre
	return node
}
