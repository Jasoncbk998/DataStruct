/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/3 6:59 下午
 **/
package main

//没有引用就等于删除,会对该对象进行回收
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
