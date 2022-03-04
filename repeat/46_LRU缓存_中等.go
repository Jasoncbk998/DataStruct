/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:49 下午
 **/
package main

type DlinkedNode struct {
	key, value int
	prev, next *DlinkedNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DlinkedNode
	head, tail *DlinkedNode
}

func initDlinkedNode(key, value int) *DlinkedNode {
	return &DlinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DlinkedNode{},
		head:     initDlinkedNode(0, 0),
		tail:     initDlinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := initDlinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) moveToHead(node *DlinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeNode(node *DlinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) addToHead(node *DlinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeTail() *DlinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
