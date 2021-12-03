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
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true

}

//hash表
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}
