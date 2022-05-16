/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/6 7:31 下午
 **/
package main

import "DataStruct/tools"

func hasCycle(head *tools.ListNode) bool {
	//头结点或者头结点的下一个节点是nil则false
	if head == nil || head.Next == nil {
		return false
	}
	// 快慢指针,快指针比满指针快一个节点,并且可以追上且不为nil则说明成环
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle_(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
