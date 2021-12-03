/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/19 8:27 下午
 **/
package main

import (
	Single_Link "DataStruct/dataStruct/LinkedLIst"
	"fmt"
)

func main() {
	list := Single_Link.NewSingleLinkList()
	list.InsertNodeBack(Single_Link.NewSingleLinkNode(1))
	list.InsertNodeBack(Single_Link.NewSingleLinkNode(2))
	list.InsertNodeBack(Single_Link.NewSingleLinkNode(3))
	list.InsertNodeBack(Single_Link.NewSingleLinkNode(4))
	fmt.Println(list)
	list.ReverseListOnIndex(2)
	fmt.Println(list)

}
