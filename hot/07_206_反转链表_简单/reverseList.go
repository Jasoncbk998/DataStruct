/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/21 11:22 上午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1,2,3,4,5
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	// 1,2,3,4,5
	curr := head
	for curr != nil {
		// 2,3,4,5
		//3,4,5
		next := curr.Next
		// 1,nil
		// 2,1,nil
		curr.Next = prev
		// 1,nil
		//2,1,nil
		prev = curr
		// 2,3,4,5
		// 3,4,5
		curr = next
	}
	return prev
}

func test(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
