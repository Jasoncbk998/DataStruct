/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/2 13:07
 **/
package main

import "DataStruct/tools"

/**
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

输入：head = [7,7,7,7], val = 7
输出：[]
*/
func removeElements(head *tools.ListNode, val int) *tools.ListNode {
	if head == nil {
		return nil
	}
	dummy := &tools.ListNode{Next: head}
	for cur := dummy; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
