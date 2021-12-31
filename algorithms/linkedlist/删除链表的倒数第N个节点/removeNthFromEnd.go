/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/3 7:04 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}
