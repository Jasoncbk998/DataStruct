/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/13 9:25 下午
 **/
package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{5, &ListNode{3, &ListNode{1, nil}}}}}}
	is := isPalindrome(l1)
	fmt.Println(is)
}

//快慢指针
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//因为判断是否是回文串,所以需要找到中间节点,通过快慢指针找到中间节点
	firstHalfEnd := endOfFirstHalf(head)
	//后半段
	secondHalfStart := reverseList(firstHalfEnd.Next)

	//头
	p1 := head
	//后半段的头
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		//后移
		p1 = p1.Next
		p2 = p2.Next
	}
	//还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
