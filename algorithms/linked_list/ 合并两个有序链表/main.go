/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 9:19 下午
 **/
package main

import "fmt"

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	lists := mergeTwoLists3(l1, l2)
	fmt.Println(lists)

}

func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			fmt.Println(l1.Val)
			//l1是一个链表
			cur.Next = l1
			l1 = l1.Next
		} else {
			fmt.Println(l2.Val)
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return ret.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		fmt.Println(l1.Val)
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}

}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	vir := &ListNode{}
	n := vir
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}
		n = n.Next
	}
	for n1 != nil {
		n.Next = n1
		n1 = n1.Next
		n = n.Next
	}
	for n2 != nil {
		n.Next = n2
		n2 = n2.Next
		n = n.Next
	}
	return vir.Next
}

func (list *ListNode) String() string {
	var listString string
	p := list.Next
	for p.Next != nil {
		listString += fmt.Sprintf("%v-->", p.Next.Val)
		p = p.Next
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}
