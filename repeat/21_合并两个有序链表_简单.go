/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/20 11:44 上午
 **/
package main

import (
	"DataStruct/tools"
	"fmt"
)

func mergeTwoLists(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {
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

func main() {
	l1 := &tools.ListNode{1, &tools.ListNode{3, nil}}
	l2 := &tools.ListNode{2, &tools.ListNode{4, nil}}
	s := mergeTwoLists(l1, l2).String()
	fmt.Println(s)
}
