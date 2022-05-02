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
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart

	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func reverseList(head *tools.ListNode) *tools.ListNode {
	var prev, cur *tools.ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *tools.ListNode) *tools.ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
