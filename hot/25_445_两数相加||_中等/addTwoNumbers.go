/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/26 2:07 下午
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]
*/
/**
1 2 3
  1 2
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := pushStack(l1)
	stack2 := pushStack(l2)

	dummyHead := &ListNode{}
	head := dummyHead
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		val := carry
		if len(stack1) > 0 {
			val = val + stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			val = val + stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		carry = val / 10
		tmp := head.Next
		head.Next = &ListNode{Val: val % 10, Next: tmp}
	}
	return dummyHead.Next
}

func pushStack(l *ListNode) []int {
	var stack []int
	for l != nil {
		stack = append(stack, l.Val)
		l = l.Next
	}
	return stack
}
