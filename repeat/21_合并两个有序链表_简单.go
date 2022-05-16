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

// 递归调用
/**
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func mergeTwoLists(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//不断递归
	if l1.Val < l2.Val {
		// l1的值小于l2时则l1.next
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

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
[1,2,4]
[1,3,4]
*/
func mergeTwoLists_(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists_(list1, list2.Next)
		return list1
	}
	list2.Next = mergeTwoLists_(list1.Next, list2)
	return list2
}
