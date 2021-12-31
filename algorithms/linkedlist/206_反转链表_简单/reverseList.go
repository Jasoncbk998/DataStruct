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
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//for l1 !=nil{
	//	fmt.Print(l1.Val," ")
	//	l1=l1.Next
	//}
	list := reverseList(l1)
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	//for list!=nil{
	//	fmt.Println(list.Val)
	//	list=l1.Next
	//}
}

// head = [1,2,3,4,5]
// [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		//2,3
		//3
		next := curr.Next
		//1,nil
		//2,1,nil
		curr.Next = prev
		// 1,nil
		//2,1,nil
		prev = curr
		// 2,3
		// 3
		curr = next
	}
	return prev
}
