/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 4:04 下午
 **/
package main

import (
	"DataStruct/dataStruct/Queue"
	"fmt"
)

func main() {
	queue := Queue.NewQueue()

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	//for i:=0;i<3;i++{
	//	deQueue := queue.DeQueue()
	//	fmt.Println(deQueue)
	//}

	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Size())
}
