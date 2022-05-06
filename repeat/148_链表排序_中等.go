/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/4/28 12:35
 **/
package main

import "DataStruct/tools"

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
归并不会,太难了
*/
func merge_node(head1, head2 *tools.ListNode) *tools.ListNode {
	dummyHead := &tools.ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2

	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
func sorts(head, tail *tools.ListNode) *tools.ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge_node(sorts(head, mid), sorts(mid, tail))
}

func sortList(head *tools.ListNode) *tools.ListNode {
	return sorts(head, nil)
}
