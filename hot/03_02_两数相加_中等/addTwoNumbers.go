/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/19 2:58 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

/**
思路很简单,熟练就好
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		//取个位
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		//取十位
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
