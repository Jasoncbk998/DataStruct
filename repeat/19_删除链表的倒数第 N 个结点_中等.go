/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/8 1:09 下午
 **/
package main

import "DataStruct/tools"

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
/**
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
func removeNthFromEnd(head *tools.ListNode, n int) *tools.ListNode {
	// 获取链表长度
	length := getLength(head)
	dummy := &tools.ListNode{0, head}
	cur := dummy
	//移动链表至倒数第n+1个元素位置
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	// 开始删除该元素 从第n+1个元素'跨过去'
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *tools.ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

func main() {

	node := &tools.ListNode{1, &tools.ListNode{2, &tools.ListNode{3, &tools.ListNode{4, &tools.ListNode{5, nil}}}}}
	removeNthFromEnd(node, 2)
}
