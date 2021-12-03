package main

import (
	"DataStruct/dataStruct/Tree/RBTree"
	"fmt"
)

func main() {
	rbtree := RBTree.NewRBTree()
	for i := 0; i < 1000000; i++ {
		rbtree.Insert(thanInt(i))
	}
	for i := 0; i < 900000; i++ {
		//rbtree
		rbtree.Delete(thanInt(i))
	}
	fmt.Println(rbtree.GetPath())
}
