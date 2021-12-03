/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 12:22 下午
 **/
package ArrayListStack

import (
	"DataStruct/dataStruct/Array/ArrayList"
)

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	myarray *ArrayList.ArrayList
	capsize int //最大范围
}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = ArrayList.NewArrayList()
	mystack.capsize = 10
	return mystack
}

func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
	mystack.capsize = 10
}
func (mystack *Stack) Size() int {
	return mystack.myarray.TheSize
}
func (mystack *Stack) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	}
	return false
}
func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		//最后一个元素
		last := mystack.myarray.DataStore[mystack.myarray.TheSize-1]
		//在array中删除这个元素
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)
		return last
	}
	return nil
}
func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *Stack) IsFull() bool {
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	}
	return false
}
