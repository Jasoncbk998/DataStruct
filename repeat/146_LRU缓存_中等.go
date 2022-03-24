/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:49 下午
 **/
package main

import "DataStruct/tools"

//lru基本数据结构
type LRUCache struct {
	head, tail *tools.Node
	Keys       map[int]*tools.Node
	Cap        int
}

//构造,进行初始化
func Construstor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*tools.Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	//如果有就先移除该元素,然后add
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &tools.Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

// Add 添加到头
func (this *LRUCache) Add(node *tools.Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *tools.Node) {
	// 头结点,就让head连接node.next
	// 并且node.next=nil 这样node节点没有指引,就会被回收,达到删除目的
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	// 1,2,3
	// 末尾结点
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	// 中间节点进行关联
	node.Prev.Next = node.Next // 待删除节点的前一位节点与待删除节点的后一位连接
	node.Next.Prev = node.Prev
}
