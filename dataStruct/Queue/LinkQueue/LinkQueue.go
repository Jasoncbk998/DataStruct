/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 4:55 下午
 **/
package LinkQueue

type Node struct {
	Data  interface{}
	PNext *Node
}

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() interface{}
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) length() int {
	pnext := qlk.front
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext //节点循环跳跃
		length++            //追加
	}
	return length //返回长度
}

func (qlk *QueueLink) Enqueue(value interface{}) {
	newnode := &Node{value, nil}
	if qlk.front == nil {
		qlk.front = newnode
		qlk.rear = newnode
	} else {
		qlk.rear.PNext = newnode
		qlk.rear = qlk.rear.PNext
	}
}

func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newnode := qlk.rear        //记录头部位置
	if qlk.front == qlk.rear { //只剩下一个
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.PNext
	}
	return newnode.Data
}
