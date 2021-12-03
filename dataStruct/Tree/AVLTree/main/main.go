/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/24 11:56 上午
 **/
package main

import (
	"DataStruct/dataStruct/Tree/AVLTree"
	"fmt"
)

func compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -2
	}
	if newB, ok = b.(int); !ok {
		return -2
	}
	if newA > newB {
		return 1
	} else if newA < newB {
		return -1
	} else {
		return 0
	}
}

func main() {
	tree, _ := AVLTree.NewAVLTree(3, compare)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	fmt.Println(tree.Getall())

}
