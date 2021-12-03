/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/19 9:06 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// head = [1,2,3,4,5]
// [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

}
