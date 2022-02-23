/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/23 11:50 上午
 **/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	// 1,2,3,4,5
	for i := 0; i < left-1; i++ {
		// 2,3,4,5
		pre = pre.Next
	}
	//2,3,4,5
	cur := pre.Next
	//开始反转 子头结点后边跟着需要反转的节点
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

func main() {
	node := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{}}}}}}

	data1 := appendData(&node)
	for k, v := range data1 {
		fmt.Println(k, v)
	}
	fmt.Println()
	between := reverseBetween(&node, 2, 3)
	data := appendData(between)
	for k, v := range data {
		fmt.Println(k, v)
	}

}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}

	return []int{head.Val}
}
