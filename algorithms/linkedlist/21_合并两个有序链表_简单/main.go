/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 9:19 下午
 **/
package main

import "fmt"

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{3, nil}}
	l2 := &ListNode{2, &ListNode{4, nil}}
	lists := mergeTwoLists(l1, l2)
	fmt.Println(lists)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func (list *ListNode) String() string {
	var listString string
	p := list.Next
	for p.Next != nil {
		listString += fmt.Sprintf("%v-->", p.Next.Val)
		p = p.Next
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}
