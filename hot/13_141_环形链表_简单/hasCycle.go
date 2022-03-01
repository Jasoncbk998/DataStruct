/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/22 12:54 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针
func hasCycle(head *ListNode) bool {
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
