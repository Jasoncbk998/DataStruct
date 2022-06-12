/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/6 4:38 下午
 **/
package main

import (
	"DataStruct/tools"
)

// head = [3,2,1]
// [1,2,3]
func reverseLists(head *tools.ListNode) *tools.ListNode {
	var prev *tools.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
