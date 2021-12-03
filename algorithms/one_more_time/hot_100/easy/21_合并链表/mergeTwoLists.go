/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 12:50 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
