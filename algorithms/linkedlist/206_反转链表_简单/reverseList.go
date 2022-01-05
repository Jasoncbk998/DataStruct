/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/19 9:06 下午
 **/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{3, &ListNode{2, &ListNode{1, nil}}}
	list := reverseList(l1)
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
}

// head = [3,2,1]
// [1,2,3]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		//2,1,nil
		//1,nil
		//nil
		next := curr.Next
		// 3,nil
		//2,3,nil
		//1,2,3,nil
		curr.Next = prev
		// 3,nil
		//2,3,nil
		//1,2,3,nil
		prev = curr
		//2,1
		//1
		//nil
		curr = next
	}
	return prev
}
