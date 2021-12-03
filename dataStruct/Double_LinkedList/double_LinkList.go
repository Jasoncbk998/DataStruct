/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 7:44 下午
 **/
package Double_LinkedList

import "fmt"

type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

// 新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}
}

func (dlist *DoubleLinkList) Getlength() int {
	return dlist.length
}

//返回第一个节点
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

//头部插入  在前边插入
func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil
		phead.next = node
		node.prev = phead
		dlist.length++
	} else {
		phead.next.prev = node
		node.next = phead.next
		phead.next = node
		node.prev = phead
		dlist.length++
	}
}

func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil
		phead.next = node
		node.prev = phead
		dlist.length++
	} else {
		for phead.next != nil {
			phead = phead.next
		}
		phead.next = node
		node.prev = phead
		dlist.length++
	}
}

//末未插入
func (dlist *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil {
		//正向循环
		listString1 += fmt.Sprintf("%v-->", phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for phead != dlist.head {
		//反向循环
		listString2 += fmt.Sprintf("<--%v", phead.value)
		phead = phead.prev
	}
	listString1 += fmt.Sprintf("nil")

	return listString1 + listString2 + "\n" //打印链表字符串
}

func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest { //循环查找
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next.next != nil {
			phead.next.next.prev = node //500100
		}
		node.next = phead.next.next //300300
		phead.next.next = node      //500100
		node.prev = phead.next      //300200
		//dlist.head=phead
		dlist.length++
		return true
	} else {
		return false
	}
}

func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest { //循环查找
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next != nil {
			phead.next.prev = node //400500
		}
		node.next = phead.next //300100
		node.prev = phead      //300200
		phead.next = node      //400500
		//dlist.head=phead
		dlist.length++
		return true
	} else {

		return false
	}
}

func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	if node == nil {
		return false
	} else {
		phead := dlist.head
		for phead.next != nil && phead.next != node {
			phead = phead.next //循环查找
		}
		if phead.next == node {
			if phead.next.next != nil {
				phead.next.next.prev = phead //设置pre
			}
			phead.next = phead.next.next //设置next

			dlist.length--
			return true
		} else {
			return false
		}

		return true

	}
}

func (dlist *DoubleLinkList) DeleteNodeAtindex(index int) bool {
	if index > dlist.length-1 || index < 0 {
		return false
	}
	phead := dlist.head
	for index > 0 {
		phead = phead.next
		index-- //计算位置
	}

	if phead.next.next != nil {
		phead.next.next.prev = phead //设置pre
	}
	phead.next = phead.next.next //设置next

	dlist.length--
	return true

}
