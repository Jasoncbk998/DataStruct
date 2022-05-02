/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/6 7:33 下午
 **/
package tools

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode() *ListNode {
	return &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
}

func (listnode *ListNode) String() string {
	var listString string
	temp := listnode
	//p := listnode.Next
	if listnode.Val == 0 {
		temp = listnode.Next
	}
	for temp != nil {
		listString += fmt.Sprintf("%v-->", temp.Val)
		temp = temp.Next
	}
	listString += fmt.Sprintf("nil")
	return listString
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Key, Val int
	// 前 ,后
	Prev, Next *Node
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return b
}
