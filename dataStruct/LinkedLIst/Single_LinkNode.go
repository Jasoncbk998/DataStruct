package Single_Link

//链表的节点
type SingleLinkNode struct {
	value interface{}     // 数据
	pNext *SingleLinkNode //节点中指向下一个节点的指针
	// 通常是 pNext是一个指针 连接到下一个节点
}

//构造一个节点
func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

//返回数据
func (node *SingleLinkNode) Value() interface{} {
	return node.value
}

//返回节点
func (node *SingleLinkNode) PNext() *SingleLinkNode {
	return node.pNext
}
