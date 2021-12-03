/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 8:55 下午
 **/
package main

import (
	"DataStruct/dataStruct/Double_LinkedList"
	"fmt"
)

func main() {

	list := Double_LinkedList.NewDoubleLinkList()
	list.InsertBack(Double_LinkedList.NewDoubleLinkNode(1))
	list.InsertBack(Double_LinkedList.NewDoubleLinkNode(2))
	list.InsertBack(Double_LinkedList.NewDoubleLinkNode(3))
	list.InsertBack(Double_LinkedList.NewDoubleLinkNode(4))
	fmt.Println(list)
}
