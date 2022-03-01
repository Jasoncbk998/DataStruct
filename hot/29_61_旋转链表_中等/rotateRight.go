/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/1 3:25 下午
 **/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
*/
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	right := rotateRight(head, 2)
	fmt.Println(right)
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	len := 0
	cur := newHead
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	if (k % len) == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	// 进行截断
	cur.Next = nil
	return res.Next
}
