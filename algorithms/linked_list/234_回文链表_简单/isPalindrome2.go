/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{5, &ListNode{3, &ListNode{1, nil}}}}}}
	is := isPalindrome2(l1)
	fmt.Println(is)
}

func isPalindrome2(head *ListNode) bool {
	node1, node2 := split(head)
	reversedNode2 := reverse(node2)
	return equals(node1, reversedNode2)
}

func split(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	node2 := slow.Next
	slow.Next = nil
	return head, node2
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func equals(head1, head2 *ListNode) bool {
	node1, node2 := head1, head2
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
