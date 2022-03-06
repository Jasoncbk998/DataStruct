/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/6 4:38 下午
 **/
package main

// head = [3,2,1]
// [1,2,3]
func reverseList_test(head *ListNode) *ListNode {
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
