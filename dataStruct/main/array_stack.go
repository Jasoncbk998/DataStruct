/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 12:41 上午
 **/
package main

import (
	"DataStruct/dataStruct/Array/ArrayList"
	ArrayListStack2 "DataStruct/dataStruct/Array/ArrayListStack"
	"DataStruct/dataStruct/Stack/StackArray"
	"fmt"
)

func Fab(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return Fab(num-1) + Fab(num-2)
	}
}

func main5() {
	//fmt.Println(Fab(5))
	//用栈模拟斐波那契额
	stack := StackArray.NewStack()
	stack.Push(5)
	last := 0
	for !stack.IsEmpty() {
		data := stack.Pop()
		if data == 1 || data == 2 {
			last += 1
		} else {
			stack.Push((data.(int) - 2))
			stack.Push((data.(int) - 1))
		}
	}
	fmt.Println(last)
}

// 递归 的思想就很像栈
// 栈pop出一个 就变成一个小的栈  在pop,栈又变小了  思想上很接近
func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

//用栈模拟递归
func main4() {
	//fmt.Println(Add(5))
	//压入数据5,模拟递归
	mystack := ArrayListStack2.NewArrayListStack()
	mystack.Push(5)
	last := 5 //保存结果
	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			//递归
			mystack.Push((data.(int) - 1))
		}
	}
	fmt.Println(last)
}

//stackArray
func main3() {
	//stack := StackArray.NewStack()
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//listStack := ArrayListStack.NewArrayListStack()
	//listStack.Push(1)
	//listStack.Push(2)
	//listStack.Push(3)
	//listStack.Push(4)
	listStackX := ArrayListStack2.NewArrayListStackX()
	listStackX.Push(1)
	listStackX.Push(2)
	listStackX.Push(3)
	listStackX.Push(4)
	for it := listStackX.Myit; it.HasNext(); {
		next, _ := it.Next("111111")
		fmt.Println(next)
	}
}
func main2() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(1)
	list.Append(1)
	for i := 0; i <= 100; i++ {
		list.Insert(list.Size(), i)
	}
	fmt.Println(list)
	fmt.Println(list.Size())
	for i := 0; i < 50; i++ {

		list.Delete(i)
	}
	fmt.Println(list)
	fmt.Println(list.Size())
}
func main1() {
	//接口的用处
	//实现接口要实现所有的方法
	//var list ArrayList.List = ArrayList.NewArrayList()
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(1)
	list.Append(0)
	list.Append(1)
	fmt.Println(list.String())
	list.Insert(list.Size()+1, 2)
	fmt.Println(list.String())

}
