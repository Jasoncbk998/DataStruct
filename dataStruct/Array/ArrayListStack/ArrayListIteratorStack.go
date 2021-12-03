/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 12:23 下午
 **/
package ArrayListStack

import (
	ArrayList "DataStruct/dataStruct/Array/ArrayList"
)

type StackArrayX interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}

type StackX struct {
	myarray *ArrayList.ArrayList
	Myit    ArrayList.Iterator
}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = ArrayList.NewArrayList()
	mystack.Myit = mystack.myarray.Iterator()
	return mystack
}
func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize
}
func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
	mystack.myarray.TheSize = 0
}
func (mystack *StackX) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.DataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)
		return last
	}
	return nil
}

func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *StackX) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	}
	return false
}

func (mystack *StackX) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}
