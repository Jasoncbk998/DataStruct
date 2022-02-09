/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/10 2:34 下午
 **/
package main

/**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
删除链表中所有指定值的结点。
*/
//很好理解,理顺思路,链表题别着急
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	temp := &ListNode{0, head}
	pp := temp
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return pp.Next
}
func test(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	temp := &ListNode{0, head}
	pp := temp
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return pp.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
