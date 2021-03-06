/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 5:00 下午
 **/
package main

//2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
//Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
//Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
/**
为了处理方法统一，可以先建立一个虚拟头结点，这个虚拟头结点的 Next 指向真正的 head，
这样 head 不需要单独处理，直接 while 循环即可。另外判断循环终止的条件不用是 p.Next ！= nil，
这样最后一位还需要额外计算，循环终止条件应该是 p != nil。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		// val是个位数 ,比如18 18%10=8
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		//取出 十位  18/10=1
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
