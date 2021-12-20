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

func hasCycle3(head *ListNode) bool {
	if head == nil {
		return false
	}
	//slow
	temp1 := head
	//fast
	temp2 := head.Next
	for temp2 != nil {
		temp1 = temp1.Next
		temp2 = temp2.Next
		if temp2 != nil {
			temp2 = temp2.Next
		} else {
			return false
		}
		// 找到环
		if temp2 == temp1 {
			return true
		}
	}
	return false
}

//快慢指针
/**
具体地，我们定义两个指针，一快一满。慢指针每次只移动一步，而快指针每次移动两步。
初始时，慢指针在位置 head，而快指针在位置 head.next。
这样一来，如果在移动的过程中，快指针反过来追上慢指针，就说明该链表为环形链表。否则快指针将到达链表尾部，该链表不为环形链表。

*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	//两个节点,一前一后,如果追上了,就是相等,则成环
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next //快指针,比慢指针快一步
		fast = fast.Next.Next
	}
	return true

}
