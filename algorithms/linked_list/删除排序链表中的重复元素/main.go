/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 9:08 下午
 **/
package main

func main() {
	//list := a.NewLinkList()
	//a.NewLinkNode(1)
	//a.NewLinkNode(1)
	//a.NewLinkNode(1)
	//a.NewLinkNode(1)
}

/**
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
返回同样按升序排列的结果链表。
输入：head = [1,1,2]
输出：[1,2]
输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	//不断顺序比较链表的元素
	for head != nil && cur.Next != nil {
		//比较相邻的元素
		if cur.Val == cur.Next.Val {
			//如果val相同,则跨过相同元素,进行指针连接
			cur.Next = cur.Next.Next
		} else {
			//如果相邻元素不等,则移动指针进行下一个相邻元素的比较
			//更新cur
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head

}

func deleteDuplicates2(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
