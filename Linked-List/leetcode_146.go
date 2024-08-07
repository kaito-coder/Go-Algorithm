package main

func main() {

}

type LinkedListNode struct {
	key   int
	value int
	next  *LinkedListNode
	prev  *LinkedListNode
}

type LinkedList struct {
	size int
	head *LinkedListNode
	tail *LinkedListNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*LinkedListNode
	linkedList *LinkedList
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*LinkedListNode)
	return LRUCache{
		cache:      cache,
		capacity:   capacity,
		linkedList: &LinkedList{},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, found := this.cache[key]; !found {
		return -1
	}
	value := this.cache[key].value

	return value
}

func (this *LRUCache) Put(key int, value int) {

}
func (l *LinkedList) removeNode(node *LinkedListNode) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.next.prev = node.prev
	}
	if node == l.head {
		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.prev
	}
	l.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
