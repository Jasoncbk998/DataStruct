package main

import (
	"fmt"
)

/**
链表排序
*/
const (
	ERROR = -21231231231231
)

var phead *LinkNode = NewLinkList()

//定义类型
type Element int64

//链表基本结构
type LinkNode struct {
	Data  Element
	pNext *LinkNode
}

type LinkNoder interface {
	Add(head *LinkNode, data Element, ishead bool)
	AddHead(head *LinkNode, data Element)
	Append(head *LinkNode, data Element)
	Show(head *LinkNode)
	Getlength(head *LinkNode) int
	Search(head *LinkNode, data Element) bool
	Getdata(head *LinkNode, id int) Element
	Delete(head *LinkNode, id int) Element
	clear(phead *LinkNode)
	Insert(phead *LinkNode, index int, data Element)
	Reverse(head *LinkNode)
	BubbleSort(head *LinkNode)
	SelectSort(head *LinkNode)
	InsertSort(head *LinkNode)
	MergeSort(head *LinkNode)
	QucikSort(head *LinkNode)
}

//创建头部节点，不放任何数据
func NewLinkList() *LinkNode {
	return &LinkNode{Data: 0, pNext: nil}
}
func Add(head *LinkNode, data Element, ishead bool) {
	if ishead {
		AddHead(head, data)
	} else {
		Append(head, data)
	}
}
func AddHead(head *LinkNode, data Element) {
	node := &LinkNode{Data: data} //新建一个节点
	node.pNext = head.pNext
	head.pNext = node
}
func Append(head *LinkNode, data Element) {
	node := &LinkNode{Data: data} //新建一个节点
	if head.pNext == nil {
		head.pNext = node
	} else {
		pcur := head
		for pcur.pNext != nil {
			pcur = pcur.pNext //循环到最后
		}
		pcur.pNext = node //尾部插入
	}
}
func show(head *LinkNode) {
	pahead := head //从头节点下一个节点开始
	for pahead != nil {
		fmt.Println(pahead.Data)
		pahead = pahead.pNext //循环
	}
	fmt.Println("show over")
}

func Getlength(head *LinkNode) int {
	pahead := head.pNext //从头节点下一个节点开始
	var length int = 0
	for pahead != nil {
		//fmt.Println(pahead.Data)
		length++
		pahead = pahead.pNext //循环
	}
	//fmt.Println("show over")
	return length
}
func Search(head *LinkNode, data Element) bool {
	pahead := head.pNext //从头节点下一个节点开始
	var isfind bool = false
	for pahead != nil {
		if pahead.Data == data {
			isfind = true
			break
		}
		pahead = pahead.pNext //循环
	}
	//fmt.Println("show over")
	return isfind
}
func Getdata(head *LinkNode, id int) Element {
	if id < 0 || id > Getlength(head) {
		return ERROR
	} else {
		pgo := head
		for i := 0; i < id; i++ {
			pgo = pgo.pNext
		}
		return pgo.Data
	}
}
func Delete(head *LinkNode, id int) Element {
	if id < 0 || id > Getlength(head) {
		return ERROR
	} else {
		pgo := head
		for i := 0; i < id; i++ {
			pgo = pgo.pNext
		}
		//要删除的位置
		data := pgo.pNext.Data

		pgo.pNext = pgo.pNext.pNext //删除

		return data

	}
}
func clear(phead *LinkNode) {
	phead = nil
}
func Insert(head *LinkNode, id int, data Element) {
	if id < 0 || id > Getlength(head) {
		return
	} else {
		pgo := head
		for i := 0; i < id; i++ {
			pgo = pgo.pNext
		}
		var node LinkNode //新节点
		node.Data = data  //插入
		node.pNext = pgo.pNext
		pgo.pNext = &node

	}
}

//反转
func reverseNode(head *LinkNode) *LinkNode {
	var pnow *LinkNode = head //第一个节点
	var pre *LinkNode = nil
	var pnext *LinkNode = nil
	var ptail *LinkNode = nil
	//head=head.pNext
	for pnow != nil {
		pnext = pnow.pNext //保存当前节点下一个节点
		if pnext == nil {
			ptail = pnow
		}
		pnow.pNext = pre
		pre = pnow
		pnow = pnext
	}
	return ptail

}
func reverseNodeList(head *LinkNode) *LinkNode {
	//如果链表为空或者链表中只有一个元素
	if head == nil || head.pNext == nil {
		return head
	} else {
		//先反转后面的链表，走到链表的末端结点
		var newhead *LinkNode = reverseNodeList(head.pNext)
		//再将当前节点设置为后面节点的后续节点
		head.pNext.pNext = head
		head.pNext = nil

		return newhead
	}
}

func main_bubble() {
	var phead *LinkNode = NewLinkList()
	Append(phead, 5)
	Append(phead, 4)
	Append(phead, 3)
	Append(phead, 1)
	AddHead(phead, 22)

	BubbleSort(phead)
	show(phead)
}

//冒泡排序
func BubbleSort(head *LinkNode) {
	for phead1 := head.pNext; phead1.pNext != nil; phead1 = phead1.pNext {
		//内层for循环,每一次循环都找到一个极大值节点
		for phead2 := head.pNext; phead2.pNext != nil; phead2 = phead2.pNext {
			// 取出两个节点进行比较,进行互换位置
			if phead2.Data < phead2.pNext.Data {
				phead2.pNext.Data, phead2.Data = phead2.Data, phead2.pNext.Data
			}
		}
	}
}

func main() {

	Append(phead, 5)
	Append(phead, 4)
	Append(phead, 3)
	Append(phead, 1)
	AddHead(phead, 22)
	selectsort(phead)
	show(phead)

}

//选择
//13579468
//9  1357    468
func selectsort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext == nil {
		return head
	} else {
		fmt.Println("select start ")
		//假定头节点为空，
		for newhead := head; newhead != nil; newhead = newhead.pNext {
			//fmt.Println(newhead.Data)
			pahead := newhead //从头节点下一个节点开始
			maxnode := newhead
			// 找到极大值
			for pahead != nil {
				if pahead.Data < maxnode.Data {
					maxnode = pahead
				}
				pahead = pahead.pNext //循环
			}
			//交换数值
			maxnode.Data, newhead.Data = newhead.Data, maxnode.Data

		}
		/*
			pahead:=head.pNext  //从头节点下一个节点开始
			maxnode:=head.pNext
			for pahead!=nil{
				if pahead.Data>maxnode.Data{
					maxnode=pahead
				}
				pahead=pahead.pNext//循环
			}
			maxnode.Data,head.pNext.Data=head.pNext.Data,maxnode.Data
		*/
		return head
	}

}

// 3   1479
//13   479
//134  79
func InsertSort(head *LinkNode) *LinkNode {

	if head == nil || head.pNext == nil {
		return head
	} else {
		var pstart *LinkNode = new(LinkNode)
		var pend *LinkNode = head
		var p *LinkNode = head.pNext
		pstart.pNext = head //存储头部的位置

		for p != nil {
			var ptmp *LinkNode = pstart.pNext
			var pre *LinkNode = pstart             //记录前一个位置，方便删除与插入
			for ptmp != p && p.Data >= ptmp.Data { //找到要插入的位置
				ptmp = ptmp.pNext
				pre = pre.pNext
			}
			if ptmp == p { //插入
				pend = p
			} else {
				pend.pNext = p.pNext
				p.pNext = ptmp
				pre.pNext = p
			}
			p = pend.pNext
			show(head)
		}

		head = pstart.pNext

		return head
	}

}
func MergeSort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext == nil {
		return head
	} else {
		var phead *LinkNode = head
		var qhead *LinkNode = head
		var pre *LinkNode = nil
		//归并位置，
		//7  1  2
		//8
		// 1,3,      5,   7,9
		for qhead != nil && qhead.pNext != nil {
			qhead = qhead.pNext.pNext //走两步
			pre = phead               //记录中间位置的上一步
			phead = phead.pNext       //中间
		}
		pre.pNext = nil //拆分两端

		var left *LinkNode = MergeSort(head)   //继续拆分左边
		var right *LinkNode = MergeSort(phead) //继续拆分右边
		return Merge(left, right)
	}
}

func Merge(left *LinkNode, right *LinkNode) *LinkNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var pres *LinkNode = new(LinkNode)
	var ptmp *LinkNode = pres //开辟节点，备份

	for left != nil && right != nil {
		if left.Data < right.Data {
			ptmp.pNext = left
			ptmp = ptmp.pNext
			left = left.pNext
		} else {
			ptmp.pNext = right
			ptmp = ptmp.pNext
			right = right.pNext
		}
	}
	if left != nil {
		ptmp.pNext = left
	} else {
		ptmp.pNext = right
	}
	ptmp = pres.pNext //回到最开始的位置

	return ptmp

}

func QuickSort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext == nil {
		return head
	} else {
		var tmphead LinkNode = LinkNode{0, head}
		qsortlist(&tmphead, head, nil)
		return head
	}

}
func qsortlist(headpre, head, tail *LinkNode) {
	if head != tail && head.pNext != tail {
		var mid *LinkNode = partition(headpre, head, tail) //递归
		qsortlist(headpre, headpre.pNext, mid)
		qsortlist(mid, mid.pNext, tail)
	}
}

// 3       8 9 2 1
//3       8 29  1
//3       1 2      89

func partition(lowpre, low, high *LinkNode) *LinkNode {
	key := low.Data //取得中间数据
	var node1 LinkNode = LinkNode{0, nil}
	var node2 LinkNode = LinkNode{0, nil}
	var small, big *LinkNode = &node1, &node2
	for i := low.pNext; i != high; i = i.pNext {
		if i.Data < key {
			small.pNext = i
			small = i
		} else {
			big.pNext = i
			big = i
		}
	}
	big.pNext = high
	small.pNext = low
	low.pNext = node2.pNext
	lowpre.pNext = node1.pNext
	return low
}

func reversrList(head *LinkNode) *LinkNode {
	cur := head
	var pre *LinkNode = nil
	for cur != nil {
		pre, cur, cur.pNext = cur, cur.pNext, pre //这句话最重要
	}
	return pre
}

/**
包含了链表的一些基本操作
*/
func main1() {
	var phead *LinkNode = NewLinkList()
	Append(phead, 1)
	Append(phead, 3)
	Append(phead, 2)
	Append(phead, 4)
	AddHead(phead, 23)
	show(phead)
	//selectsort(phead)
	//phead=QuickSort(phead)
	//phead=reverseNode(phead)
	//phead=reverseNode(phead)
	//phead=reverseNodeList(phead)
	//phead=MergeSort(phead)
	//phead=InsertSort(phead)
	phead = reversrList(phead)
	fmt.Println("----------------")
	//fmt.Println(phead.Data,phead.pNext.Data,phead.pNext.pNext.Data)
	fmt.Println("----------------")
	//fmt.Println(phead.Data)
	show(phead)
}

func main2() {
	var phead *LinkNode = NewLinkList()
	Append(phead, 1)
	Append(phead, 2)
	Append(phead, 3)

	show(phead)
	//fmt.Println(Delete(phead,2))
	//fmt.Println(Getlength(phead))
	//fmt.Println(Search(phead,4))
	//fmt.Println("getdata",Getdata(phead,3))
	var newnode *LinkNode = phead
	phead = reverseNode(newnode.pNext)
	//fmt.Println(phead.Data,phead.pNext.Data,phead.pNext.pNext.Data)
	show(phead)
}
