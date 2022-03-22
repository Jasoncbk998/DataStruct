/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/22 2:14 下午
 **/
package main

import (
	"DataStruct/tools"
)

/**
判断链表是否有环，不能使用额外的空间。如果有环，输出环的起点指针，如果没有环，则输出空。

给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
*/
func detectCycle(head *tools.ListNode) *tools.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	f, s := head, head
	for f != nil {
		s = s.Next
		if f.Next != nil {
			f = f.Next.Next
		} else {
			return nil
		}
		if f == s {
			t := head
			for t != s {
				s = s.Next
				t = t.Next
			}
			return s
		}

	}
	return nil
}
