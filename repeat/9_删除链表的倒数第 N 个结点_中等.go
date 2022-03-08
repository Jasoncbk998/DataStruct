/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/8 1:09 下午
 **/
package main

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
/**
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
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

func main() {

	node := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	removeNthFromEnd(node, 2)

}
