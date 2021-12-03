/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 5:18 下午
 **/
package LinkStack

// 这是一个节点
type Node struct {
	Data  interface{}
	PNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n == nil {
		return true
	} else {
		return false
	}
}
