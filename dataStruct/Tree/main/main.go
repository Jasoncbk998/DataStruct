/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/21 6:51 下午
 **/
package main

import (
	"DataStruct/dataStruct/Tree/BinaryTree"
	"fmt"
)

func main() {
	bst := BinaryTree.NewBinaryTree()
	node1 := &BinaryTree.Node{4, nil, nil}
	node2 := &BinaryTree.Node{2, nil, nil}
	node3 := &BinaryTree.Node{6, nil, nil}
	node4 := &BinaryTree.Node{1, nil, nil}
	node5 := &BinaryTree.Node{3, nil, nil}
	node6 := &BinaryTree.Node{5, nil, nil}
	node7 := &BinaryTree.Node{7, nil, nil}

	bst.Root = node1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	bst.Size = 7

	bst.RemoveMin()

}
func main1() {
	tree := BinaryTree.NewBinaryTree()
	/**
	      4
	   2       6
	1   3   5   7
	*/
	//插入一棵树
	tree.Add(4)
	tree.Add(2)
	tree.Add(6)
	tree.Add(1)
	tree.Add(3)
	tree.Add(5)
	tree.Add(7)

	/**
	中序遍历:  123 4 567
	前序遍历:  4 213  657
	后序遍历:  132 576 4
	*/
	tree.RemoveMin()
	tree.InOrder()
	fmt.Println("--------")
	tree.PreOrder()
	fmt.Println("--------")
	tree.PostOrder()

	s := tree.String()
	fmt.Println(s)

	tree.Levelshow()

}
