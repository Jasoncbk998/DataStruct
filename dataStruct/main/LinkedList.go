/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 2:43 下午
 **/
package main

import (
	Single_Link "DataStruct/dataStruct/LinkedLIst"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	node := Single_Link.NewSingleLinkList()
	node.InsertNodeBack(Single_Link.NewSingleLinkNode(1))
	node.InsertNodeBack(Single_Link.NewSingleLinkNode(2))
	node.InsertNodeBack(Single_Link.NewSingleLinkNode(3))
	node.InsertNodeBack(Single_Link.NewSingleLinkNode(4))
	node.InsertNodeValueBack(2, Single_Link.NewSingleLinkNode(4))
	fmt.Println(node)

}
func main2_() {
	path := "SearchData/houdao/1_1.txt"
	list := Single_Link.NewSingleLinkList()
	sqlfile, _ := os.Open(path)    //打开文件
	i := 0                         //统计行数
	br := bufio.NewReader(sqlfile) //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}
		linestr := string(line) //转化为字符串
		nodestr := Single_Link.NewSingleLinkNode(linestr)
		list.InsertNodeFront(nodestr) //数据插入链表
		//fmt.Println(linestr)
		i++
	}
	fmt.Println(i, "内存载入完成")

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime := time.Now()
		list.FindString(inputstr)
		fmt.Println("本次查询用了", time.Since(starttime))
	}

}
func main1_() {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node2 := Single_Link.NewSingleLinkNode(2)
	node3 := Single_Link.NewSingleLinkNode(3)
	list.InsertNodeFront(node1)
	//fmt.Println(list)
	list.InsertNodeFront(node2)
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)

	list.Deleteatindex(1)
	fmt.Println(list)

}
