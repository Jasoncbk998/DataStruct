/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 7:03 下午
 **/
package main

type node struct {
	data  int
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func NewBST() *BST { // 创建树根
	bst := &BST{}
	bst.root = nil
	bst.size = 0
	return bst
}

type BSTfunc interface {
	Getsize()
	Add()
	AddNode()
	Contains()  //包含
	BinSearch() // 二分查找
	PreOrder()  // 前序遍历
	InOrder()   // 中序遍历
	PostOrder() //后序遍历
	Min()       //极小值
	Max()       // 极大值
	Remove()    // 移除节点
	makeTree()  // 构造树
	GetDepth()  // 树的深度
}
