/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/22 1:56 下午
 **/
package main

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	//快慢指针找到中间节点
	pre, slow, fast := (*ListNode)(nil), head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//slow就是中间节点,slow.next就相当于删除了
	pre.Next = slow.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
