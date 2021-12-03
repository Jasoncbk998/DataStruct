/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 4:39 下午
 **/
package main

import (
	"DataStruct/dataStruct/Queue/CricleQueue"
	"fmt"
)

func main() {
	var cq CricleQueue.CricleQueue
	CricleQueue.InitQueue(&cq)
	CricleQueue.EnQueue(&cq, 1)
	CricleQueue.EnQueue(&cq, 2)
	CricleQueue.EnQueue(&cq, 3)
	CricleQueue.EnQueue(&cq, 4)
	CricleQueue.EnQueue(&cq, 5)
	CricleQueue.EnQueue(&cq, 6)
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
	fmt.Println(CricleQueue.DeQueue(&cq))
}
