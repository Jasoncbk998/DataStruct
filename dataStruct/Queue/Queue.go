/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 3:34 下午
 **/
package Queue

type MyQueue interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //是否为空
	EnQueue(data interface{}) //入队列
	DeQueue() interface{}     //出队列
	Clear()                   //清空
}

type Queue struct {
	dataStore []interface{} //队列中数据的存储
	theSize   int           //队列的大小
}

//构造
func NewQueue() *Queue {
	myqueue := new(Queue)
	myqueue.dataStore = make([]interface{}, 0)
	myqueue.theSize = 0
	return myqueue
}

func (myq *Queue) Size() int {
	return myq.theSize
}

//取出第一个
func (myq *Queue) Front() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[0]
}

//取出最后一个
func (myq *Queue) End() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[myq.Size()-1]
}

func (myq *Queue) IsEmpty() bool {
	return myq.theSize == 0
}
func (myq *Queue) EnQueue(data interface{}) {
	myq.dataStore = append(myq.dataStore, data) //入队
	myq.theSize++
}

func (myq *Queue) DeQueue() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	//第一个元素出队列
	data := myq.dataStore[0]
	if myq.Size() > 1 {
		myq.dataStore = myq.dataStore[1:myq.Size()]
	} else {
		myq.dataStore = make([]interface{}, 0)
	}
	myq.theSize--
	return data
}

func (myq *Queue) Clear() {
	myq.dataStore = make([]interface{}, 0)
	myq.theSize = 0
}
