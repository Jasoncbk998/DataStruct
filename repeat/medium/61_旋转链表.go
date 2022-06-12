/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/5 14:07
 **/
package main

/**
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1 // 链表的长度
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	// 算出往前移动的位数
	add := n - k%n
	if add == n {
		return head
	}
	// 这里其实是成环,因为前边已经移动到了最后一个节点 然后在拼接头结点成环
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
