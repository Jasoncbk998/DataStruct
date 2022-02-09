/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/9 8:55 下午
 **/
package _669_合并两个链表_中等

/**
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个
请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。

考查链表的基本操作。此题注意 a == b 的情况。
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	n := list1
	var startRef, endRef *ListNode
	for i := 0; i <= b; i++ {
		if i == a-1 {
			startRef = n
		}
		if i == b {
			endRef = n
		}
		n = n.Next
	}
	startRef.Next = list2
	n = list2
	for n.Next != nil {
		n = n.Next
	}
	n.Next = endRef.Next
	return list1

}

type ListNode struct {
	Val  int
	Next *ListNode
}
