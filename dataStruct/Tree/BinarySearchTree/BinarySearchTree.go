/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/25 4:08 下午
 **/
package BinarySearchTree

type Node struct {
	Data   int   //树的数据
	Parent *Node //父节点
	Left   *Node //左节点
	Right  *Node //右节点
}
type BinarySearchTree struct {
	Root *Node // 根节点
	Size int   //数据的数量
}

func NewBinarySearchTree() *BinarySearchTree {
	bst := &BinarySearchTree{}
	bst.Size = 0
	bst.Root = nil
	return bst
}

type BinaryTreeInfo interface {
	root()
	left(node interface{})
	right(node interface{})
	string(node interface{})
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.Size == 0
}

//清空树
func (bst *BinarySearchTree) Clear() {
	bst.Size = 0
	bst.Root = nil
}

func (bst *BinarySearchTree) Add(data int) {
	bst.Root = bst.add(bst.Root)
}
func (bst *BinarySearchTree) add(node *Node) *Node {

	if node == nil {
		bst.Size++
		return &Node{Data: node.Data, Parent: nil, Left: nil, Right: nil}
	}
	return nil
}
