/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/21 3:30 下午
 **/
package main

import (
	"DataStruct/tools"
)

/**
找到 2 个链表的交叉点。
*/
func getIntersectionNode(headA, headB *tools.ListNode) *tools.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	//不断比较节点
	//A 1,1,2,3,4,5
	//B 6,4,5
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		//fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
