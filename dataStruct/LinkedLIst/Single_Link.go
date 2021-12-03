package Single_Link

import (
	"fmt"
	"strings"
)

//单链表的接口，
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode        //抓取头部节点
	InsertNodeFront(node *SingleLinkNode) //头部插入
	InsertNodeBack(node *SingleLinkNode)  //尾部插入

	//在一个节点之前插入或者一个节点之后插入
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool
	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool

	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定位置的节点

	DeleteNode(dest *SingleLinkNode) bool //删除一个节点
	Deleteatindex(index int)              //删除指定位置的节点

	String() string //返回链表字符串

}

// SingleLinkList 链表的结构
type SingleLinkList struct {
	head   *SingleLinkNode //链表的头指针
	length int             //链表的长度
}

// NewSingleLinkList 创建链表
func NewSingleLinkList() *SingleLinkList {
	//头结点,nil 空的头结点
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{head, 0}
}

//返回第一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext // 返回的是一个指针
}

//头部插入
func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.head == nil {
		list.head.pNext = node //对头结点赋值
		node.pNext = nil       // 对node节点下一个进行nil
		list.length++          //插入节点，数据追加
	} else {
		//头节点备份
		bak := list.head
		//node的下一个节点指定为原头结点
		node.pNext = bak.pNext
		//把bak的节点指针换成node,完成了插入到首位
		bak.pNext = node
		list.length++ //插入节点，数据追加
	}
}

//尾部插入
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.head == nil {
		// 如果head是nil,说明链表只有head
		list.head.pNext = node // 所以直接head指向下一个节点的指针直接复制为node
		node.pNext = nil       // node的下一个节点的指针赋值为nil,代表下一个节点是nil
		list.length++          //插入节点，数据追加
	} else {
		// head下边有节点
		bak := list.head // 拷贝一个副本
		for bak.pNext != nil {
			bak = bak.pNext //循环到最后
		}
		//此时经过循环后,bak是最后一个节点,将bak指向下一个节点的指针指向node节点
		bak.pNext = node //这个bak的内存地址就是原链表的地址,打断点看到的,所以直接加到这里就行了
		//bak只是一个局部变量,执行完一个方法后自动释放
		list.length++ //插入节点，数据追加
	}
}
func (list *SingleLinkList) String() string {
	var listString string
	p := list.head //头部节点
	for p.pNext != nil {
		listString += fmt.Sprintf("%v-->", p.pNext.value)
		p = p.pNext //循环
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}

func (list *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head // 找到头节点
	isfind := false    //是否找到数据
	//判断是不是只有一个头节点,如果不是则if判断
	for phead.pNext != nil {
		if phead.value == dest { //找到
			isfind = true
			break
		}
		//指针进行后移
		phead = phead.pNext
	}
	if isfind {
		//尾部插入
		node.pNext = phead.pNext // 因为phead插入之前后边有节点,所以在这里把他的后续节点指针赋值给nodeΩ
		phead.pNext = node       // 将phead的后续节点指针指向node节点
		list.length++
		return true
	} else {
		return false
	}

}
func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	//头结点
	phead := list.head
	isfind := false //是否找到数据
	//判断尾节点
	for phead.pNext != nil {
		if phead.pNext.value == dest { //找到
			isfind = true
			break
		}
		//节点移动
		phead = phead.pNext
	}
	if isfind {
		//移动指针.将phead的后续节点指针移动到node的后续节点指针
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return isfind
	} else {
		return isfind
	}

}

func (list *SingleLinkList) FindString(data string) {
	phead := list.head.pNext //指定头部
	for phead.pNext != nil { //循环所有数据
		if strings.Contains(phead.value.(string), data) { //包含
			fmt.Println(phead.value)
		}
		phead = phead.pNext
	}
}

func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext //向后循环
			index--             //向后循环过程
		}
		return phead

	}
}

// DeleteNode  删除节点
// 移动指针,覆盖值
/**
node是待删除节点
node.val  = node.next.val // 值覆盖
node.next =node.next.next // 指针移动
*/
func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	phead := list.head
	for phead.pNext != nil && phead.pNext != dest {
		phead = phead.pNext //循环下去
	}
	if phead.pNext == dest {
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	} else {
		return false
	}
}
func (list *SingleLinkList) Deleteatindex(index int) {
	if index > list.length-1 || index < 0 {
		return
	} else {
		phead := list.head
		for index > 0 {
			phead = phead.pNext //向后循环
			index--             //向后循环过程
		}
		phead.pNext = phead.pNext.pNext
		list.length--
		return
	}
}

//1>2>3>4>5
//2>3>4>5
//--   --  /
//- -
//求链表分数位置

// GetMid 快慢指针
//快指针走到最后,那么慢指针就刚好到中间
//比如1/5 则快指针走5步,慢指针走一步
func (list *SingleLinkList) GetMid() *SingleLinkNode {
	if list.head.pNext == nil {
		return nil
	} else {
		//快慢指针
		phead1 := list.head
		phead2 := list.head
		for phead2 != nil && phead2.pNext != nil {
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext
		}
		return phead1 //中间节点
	}
}

// ReverseListOnIndex 指定位置反转链表,指定元素位置  从0开始计位
func (list *SingleLinkList) ReverseListOnIndex(index int) {
	if index > list.length-1 || index < 0 {
		return
	}
	var cur *SingleLinkNode
	for index > 0 {
		cur = list.head.pNext
		index--
	}
	var pre *SingleLinkNode
	for cur.pNext != nil {
		curNext := cur.pNext
		cur.pNext = pre
		pre = cur
		cur = curNext
	}
	list.head.pNext.pNext = nil
	list.head.pNext = pre
}

// ReverseList 链表反转
//  1 -> 2 -> 3 -> 4
func (list *SingleLinkList) ReverseList() {
	if list.head == nil || list.head.pNext == nil {
		return //链表为空或者链表只有一个节点
	} else {
		var pre *SingleLinkNode                   //前面节点
		var cur *SingleLinkNode = list.head.pNext //当前节点
		//不断的循环,当前节点
		for cur != nil {
			//curNext 保存每次进入循环的子链表
			//每一次进入,curNext都会从当前节点后移一个元素
			// 1 2 3 4 5
			// 2 3 4 5
			// 3 4 5
			// ...
			curNext := cur.pNext
			//当前节点每次循环的下一个节点都复制为pre,也就是上一次循环拿出的头部节点
			//倒序的节点都是存储在pre中
			//每次都是取出当前的首节点后边追加上逆序的节点
			// 1
			// 2 1
			// 3 2 1
			// 4 3 2 1
			cur.pNext = pre
			// 因为cur.next=pre 把pre追加到cur后边
			// 所以需要把新的逆序子链表赋值给pre,等待下一次循环
			pre = cur
			// 把这次循环保存好的curNext 赋值给cur进入下一次循环
			cur = curNext //持续推进
		}
		//头结点下一个的下一个节点置为ni
		list.head.pNext.pNext = nil
		//头结点下一个节点置为pre
		list.head.pNext = pre
	}

}
