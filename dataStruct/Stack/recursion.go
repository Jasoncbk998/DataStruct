/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 7:09 下午
 **/
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归

type Node struct {
	data interface{}
	next *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(value int)
	Pop() (int, error)
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.next == nil
}

func (n *Node) Push(value interface{}) {
	newnode := &Node{data: value}
	newnode.next = n.next
	n.next = newnode
}

func (n *Node) Pop() (interface{}, error) {
	if n.IsEmpty() == true {
		return -1, errors.New("bug")
	}
	value := n.next.data
	n.next = n.next.next
	return value, nil
}

func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.next != nil {
		pnext = pnext.next
		length++
	}
	return length
}

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

// 递归转栈
//栈模拟递归
//通过不断弹出和压栈,模拟了递归
func main1() {
	stack := NewStack()
	stack.Push(10)
	last := 0
	for !stack.IsEmpty() {
		data, err := stack.Pop()
		if err != nil {
			break
		}
		//
		if data == 1 || data == 2 {
			last += 1
		} else {
			stack.Push(data.(int) - 1)
			stack.Push(data.(int) - 2)
		}
	}
	fmt.Println(last)
}

// 遍历文件夹,使用栈遍历,只要遇到文件夹,就继续压栈,遇到的是文件就直接弹栈append到string中
func main() {
	files := []string{}
	path := "dataStruct"
	mystack := NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path, err := mystack.Pop()
		if err != nil {
			break
		}
		files = append(files, path.(string))
		read, err := ioutil.ReadDir(path.(string))
		if err != nil {
			break
		}
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string) + "/" + fi.Name()
				mystack.Push(fulldir)
			} else {
				fullname := path.(string) + "/" + fi.Name() //处理文件
				files = append(files, fullname)             //插入列表
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
