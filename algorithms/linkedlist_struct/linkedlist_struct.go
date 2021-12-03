/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 9:34 下午
 **/
package linkedlist_struct

import "fmt"

//节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//链表
type LinkList struct {
	head   *ListNode //链表的头指针
	length int       //链表的长度
}

//构造一个节点
func NewLinkNode(val int) *ListNode {
	return &ListNode{val, nil}
}

//new 链表
func NewLinkList() *LinkList {
	head := NewLinkNode(0)
	return &LinkList{head, 0}
}

func (list *ListNode) String() string {
	var listString string
	p := list.Next //头部节点
	for p.Next != nil {
		listString += fmt.Sprintf("%v-->", p.Next.Val)
		p = p.Next //循环
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}
