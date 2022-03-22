/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/22 7:02 下午
 **/
package main

import "DataStruct/tools"

/**
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
*/
func reverseBetween(head *tools.ListNode, left int, right int) *tools.ListNode {
	if head == nil || left >= right {
		return head
	}
	newHead := &tools.ListNode{0, head}
	pre := newHead
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}
