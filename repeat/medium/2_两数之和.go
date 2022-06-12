/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/25 21:56
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
// 递归,链表,数学
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	n1, n2, carry := 0, 0, 0
	current := result
	for l1 != nil || l2 != nil || carry != 0 {

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return result.Next
}
