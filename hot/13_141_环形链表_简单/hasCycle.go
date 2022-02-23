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
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
