/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/22 7:30 下午
 **/
package main

import "DataStruct/tools"

func isPalindromeListNode(head *tools.ListNode) bool {
	if head == nil {
		return true
	}
	//找到中间位置
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//反转
	var prev, cur *tools.ListNode = nil, slow
	for cur != nil {
		next := cur.Next
		prev = cur
		cur = next
	}
	// slow是后半部分链表
	anthor := prev
	p, q := head, anthor
	for p != nil && p != anthor {
		if p.Val != q.Val {
			return false
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return true
}
