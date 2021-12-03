/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/18 7:40 下午
 **/
package Double_LinkedList

type DoubleLinkNode struct {
	value interface{}
	prev  *DoubleLinkNode
	next  *DoubleLinkNode
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}

func (node *DoubleLinkNode) Value() interface{} {
	return node.value
}
func (node *DoubleLinkNode) Prev() *DoubleLinkNode {
	return node.prev
}
func (node *DoubleLinkNode) Pnext() *DoubleLinkNode {
	return node.next
}
