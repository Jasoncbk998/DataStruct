/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 10:08 上午
 **/
package main

//判断链表中是否有环。

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针
/**
具体地，我们定义两个指针，一快一满。慢指针每次只移动一步，而快指针每次移动两步。
初始时，慢指针在位置 head，而快指针在位置 head.next。
这样一来，如果在移动的过程中，快指针反过来追上慢指针，就说明该链表为环形链表。否则快指针将到达链表尾部，该链表不为环形链表。
*/
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			//相等则成环
			return true
		}
	}
	//反之亦然
	return false
}
