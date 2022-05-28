/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/26 22:54
 **/
package main

type DlinkedNode struct {
	key, value int
	prev, next *DlinkedNode
}

type LRUCache struct {
	size, capactiy int
	cache          map[int]*DlinkedNode
	head, tail     *DlinkedNode
}

func initDlinkedNode(key, value int) *DlinkedNode {
	return &DlinkedNode{key: key, value: value}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DlinkedNode{},
		head:     initDlinkedNode(0, 0),
		tail:     initDlinkedNode(0, 0),
		capactiy: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDlinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capactiy {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *DlinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DlinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DlinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeTail() *DlinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
