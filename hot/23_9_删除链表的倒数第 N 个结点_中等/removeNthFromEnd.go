/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/24 12:46 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
/**
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}
