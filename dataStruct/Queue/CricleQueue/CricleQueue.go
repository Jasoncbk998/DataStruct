/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 3:41 下午
 **/
package CricleQueue

import "errors"

const QueueSize = 5 //最多存储QueueSize-1个数据
type CricleQueue struct {
	data  [QueueSize]interface{} //存储数据的结构
	front int                    //头部的位置
	rear  int                    // 尾部的位置
}

func InitQueue(q *CricleQueue) { //初始化，头部尾部重合，为空
	q.front = 0
	q.rear = 0
}

//进队列
func EnQueue(q *CricleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已经满了")
	}
	q.data[q.rear] = data
	q.rear = (q.rear + 1) % QueueSize //循环 101 , 1
	return nil
}

//出队列
func DeQueue(q *CricleQueue) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("队列为空")
	}
	res := q.data[q.front]              //取出第一个
	q.data[q.front] = 0                 // 清空
	q.front = (q.front + 1) % QueueSize //小于100+1  (100+1)%100 = 1  回到1的位置
	return res, nil
}
func Queuelength(q *CricleQueue) int { //队列长度
	//比如一直加数据,头部位置为80  尾部位置为10 相减为负数 这个时候需要QueueSize来做平衡
	return (q.rear - q.front + QueueSize) % QueueSize
}
