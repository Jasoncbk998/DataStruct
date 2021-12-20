/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 9:02 下午
 **/
package BinaryTree

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	Data  int   //树的数据
	Left  *Node //左节点
	Right *Node //右节点
}

type BinaryTree struct {
	Root *Node // 根节点
	Size int   //数据的数量
}

// NewBinaryTree 返回一个二叉树的指针地址
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{}
	bst.Size = 0
	bst.Root = nil
	return bst
}

//获取二叉树大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

//判断是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

//根节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

func (bst *BinaryTree) add(n *Node, data int) *Node {
	if n == nil {
		bst.Size++
		return &Node{data, nil, nil}
	}
	//插入数据
	if data < n.Data {
		n.Left = bst.add(n.Left, data)
	} else if data > n.Data {
		n.Right = bst.add(n.Right, data)
	}
	return n
}

//判断数据是否存在
func (bst *BinaryTree) Isin(data int) bool {
	return bst.isin(bst.Root, data)
}

func (bst *BinaryTree) isin(n *Node, data int) bool {
	if n == nil {
		return false //树是空树 找不到
	}
	if data == n.Data {
		return true
	} else if data < n.Data {
		//递归
		return bst.isin(n.Left, data)
	} else {
		return bst.isin(n.Right, data)
	}
}

func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmax(bst.Root).Data
}

func (bst *BinaryTree) findmax(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findmax(n.Right)
	}
}

func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmin(bst.Root).Data
}

func (bst *BinaryTree) findmin(n *Node) *Node {
	if n.Left == nil {
		return n
	} else {
		return bst.findmin(n.Left)
	}
}

//前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preOrder(bst.Root)
}

func (bst *BinaryTree) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)   //父节点
	bst.preOrder(node.Left)  //左节点
	bst.preOrder(node.Right) //右节点
}

//前序遍历,非递归
func (bst *BinaryTree) PreOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0)
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			res = append(res, mybst.Data) //压入数据
			mystack.PushBack(mybst)       //首先左边压入栈中
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()     // 依次取出节点
			mybst = v.Value.(*Node) //实例化
			mybst = mybst.Right     //追加
			mystack.Remove(v)       //删除
		}
	}
	return res
}

func (bst *BinaryTree) preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	bst.preorder(node.Left)
	bst.preorder(node.Right)
}

//中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inorder(bst.Root)

}

func (bst *BinaryTree) InOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组,容纳中序的数据
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst) // 首先压入左边的栈中
			mybst = mybst.Left
		}
	}
	if mystack.Len() != 0 {
		v := mystack.Back()           //挨个取出节点
		mybst = v.Value.(*Node)       //实例化
		res = append(res, mybst.Data) //压入数据
		mybst = mybst.Right           //追加
		mystack.Remove(v)             //删除
	}
	return res
}

//压入数据
//for  栈不为空
// if  else
func (bst *BinaryTree) inorder(node *Node) {
	if node == nil {
		return
	}
	bst.inorder(node.Left)
	fmt.Println(node.Data)
	bst.inorder(node.Right)
}

//后序遍历
func (bst *BinaryTree) PostOrder() {
	bst.postorder(bst.Root)
}

func (bst *BinaryTree) PostOrderNoRecursion() []int {
	mybst := bst.Root //备份
	mystack := list.New()
	res := make([]int, 0) //生成数组,容纳中序的数据
	var PreVisited *Node
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst)
			mybst = mybst.Left //左边
		}
		v := mystack.Back() //取出节点
		top := v.Value.(*Node)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && PreVisited == top.Left) || (PreVisited == top.Right) {
			res = append(res, top.Data) //压入数据
			PreVisited = top            //记录上一个节点
			mystack.Remove(v)           //处理完了在栈中删除
		} else {
			mybst = top.Right //右边循环
		}
	}
	return res
}

func (bst *BinaryTree) postorder(node *Node) {
	if node == nil {
		return
	}
	bst.postorder(node.Left)
	bst.postorder(node.Right)
	fmt.Println(node.Data)
}

func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer                     //保存字符串
	bst.GenerateBSTstring(bst.Root, 0, &buffer) //调用函数实现遍历
	return buffer.String()
}

func (bst *BinaryTree) GenerateBSTstring(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		return
	}
	bst.GenerateBSTstring(node.Left, depth+1, buffer)
	buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Data) + "\n")
	bst.GenerateBSTstring(node.Right, depth+1, buffer)
}

func (bst *BinaryTree) GenerateDepthstring(depth int) string {
	var buffer bytes.Buffer //保存字符串
	for i := 0; i < depth; i++ {
		buffer.WriteString("--") //深度为0，1-- 2----
	}
	return buffer.String()
}

//删除最小
func (bst *BinaryTree) RemoveMin() int {
	//找到最小值所在的子树里边
	ret := bst.FindMin()
	// 执行删除逻辑
	bst.Root = bst.removemin(bst.Root)
	return ret
}

//    2
//  1   3
func (bst *BinaryTree) removemin(n *Node) *Node {
	//最小值通常在左边,先判断左节点是否为nil
	if n.Left == nil {
		//删除
		// 左节点是nil,则删除他的右节点  左节点没有,就返回有节点
		rightNode := n.Right //备份右边节点
		bst.Size--           //删除
		return rightNode     //原来的节点,不要了,返回这个新的最小节点
	}
	//注意这里是递归,每一次return 都返回给上一次进来的递归函数,一层一层的推进
	//这一步就是相当于删除的操作
	//因为这个left返回后是直接复制给n  .left的,相当于root节点,改变了地址值,所以相当于删除
	n.Left = bst.removemin(n.Left) //原来的bst节点不要了 返回新的n节点 ,这个逻辑就是删除逻辑
	return n
}

//删除最大
func (bst *BinaryTree) RemoveMax() int {
	ret := bst.FindMax()
	bst.Root = bst.removemax(bst.Root)
	return ret
}

func (bst *BinaryTree) removemax(n *Node) *Node {
	if n.Right == nil {
		leftNode := n.Left
		bst.Size--
		return leftNode
	}
	n.Right = bst.removemax(n.Right)
	return n
}

func (bst *BinaryTree) Remove(data int) {
	bst.Root = bst.remove(bst.Root, data) //删除数据
}

func (bst *BinaryTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil //节点为空，不需要干活
	}
	if data < n.Data {
		n.Left = bst.remove(n.Left, data) //递归左边
		return n
	} else if data > n.Data {
		n.Right = bst.remove(n.Right, data) //递归右边
		return n
	} else {
		//处理左边为空
		if n.Left == nil {
			rightNode := n.Right //备份右边节点
			n.Right = nil
			bst.Size-- //删除
			return rightNode
		}
		//处理右边为空
		if n.Right == nil {
			leftNode := n.Left //备份右边节点
			n.Left = nil       //处理节点返回
			bst.Size--         //删除
			return leftNode
		}
		//左右节点都不为空
		oknode := bst.findmin(n.Right)        //找出比我小的节点顶替我
		oknode.Right = bst.removemin(n.Right) //6，7
		oknode.Left = n.Left                  //删除
		n.Left = nil                          //删除的清空
		n.Right = nil
		return oknode //实现删除
	}
}

// Levelshow 层级show 广度遍历
func (bst *BinaryTree) Levelshow() {
	bst.levelshow(bst.Root)
}

func (bst *BinaryTree) levelshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Front() //前面取出数据
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

//广度遍历
func (bst *BinaryTree) Stackshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Back() //前面取出数据 ,此时是栈
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) FindlowerstAncestor(root *Node, nodea *Node, nodeb *Node) *Node {
	if root == nil {
		return nil
	}
	if root == nodea || root == nodeb {
		return root //有一个节点是根节点，
	}
	left := bst.FindlowerstAncestor(root.Left, nodea, nodeb)   //递归查找
	right := bst.FindlowerstAncestor(root.Right, nodea, nodeb) //递归查找
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	lengthleft := bst.GetDepth(root.Left)
	rightlength := bst.GetDepth(root.Right)
	if lengthleft > rightlength {
		return lengthleft + 1
	} else {
		return rightlength + 1
	}
}
