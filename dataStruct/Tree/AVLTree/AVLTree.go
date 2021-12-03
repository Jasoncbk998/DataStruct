/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/24 11:00 上午
 **/
package AVLTree

import (
	"errors"
	"fmt"
)

type AVLnode struct {
	Data   interface{}
	Left   *AVLnode
	Right  *AVLnode
	height int //高度
}

//函数指针类型
type comparator func(a, b interface{}) int

//函数指针

var compare comparator

//比较大小
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func NewNode(data interface{}) *AVLnode {
	node := new(AVLnode) //新建节点初始化
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.height = 1
	return node
}

//新建avltree
func NewAVLTree(data interface{}, myfunc comparator) (*AVLnode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myfunc
	return NewNode(data), nil
}

//获取所有的值
func (avlnode *AVLnode) Getall() []interface{} {
	values := make([]interface{}, 0)
	return AddValues(values, avlnode)
}

//获取所有的值
func AddValues(values []interface{}, avlnode *AVLnode) []interface{} {
	if avlnode != nil {
		//应用递归
		values = AddValues(values, avlnode.Left)
		values = append(values, avlnode.Data)
		fmt.Println(avlnode.Data, avlnode.height)
		values = AddValues(values, avlnode.Right)
	}
	return values
}

//左旋,逆时针
func (avlnode *AVLnode) LeftRotate() *AVLnode {
	//开始左旋
	headnode := avlnode.Right
	avlnode.Right = headnode.Left
	headnode.Left = avlnode
	//更新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode
}

//右旋,顺时针
func (avlnode *AVLnode) RightRotate() *AVLnode {
	headnode := avlnode.Left //保存左边节点
	//改变指针,开始旋转
	avlnode.Left = headnode.Right
	headnode.Right = avlnode
	//更新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode
}

//两次左旋,两次右旋

//先左旋再右旋
func (avlnode *AVLnode) LeftThenRightRotate() *AVLnode {
	sonheadnode := avlnode.Left.LeftRotate()
	avlnode.Left = sonheadnode
	return avlnode.RightRotate()
}

//先右旋,在左旋
func (avlnode *AVLnode) RightThenLeftRotate() *AVLnode {
	sonheadnode := avlnode.Right.RightRotate()
	avlnode.Right = sonheadnode
	return avlnode.LeftRotate()
}

//自动处理不平衡,差距为1 为平衡  差距为2 为不平衡
func (avlnode *AVLnode) adjust() *AVLnode {
	if avlnode.Right.GetHeight()-avlnode.Left.GetHeight() == 2 {
		if avlnode.Right.Right.GetHeight() > avlnode.Right.Left.GetHeight() {
			avlnode = avlnode.LeftRotate()
		} else {
			avlnode = avlnode.RightThenLeftRotate()
		}
	} else if avlnode.Left.GetHeight()-avlnode.Right.GetHeight() == 2 {
		if avlnode.Left.Left.GetHeight() > avlnode.Left.Right.GetHeight() {
			avlnode = avlnode.RightRotate()
		} else {
			avlnode = avlnode.LeftThenRightRotate()
		}
	}
	return avlnode
}

//insert
func (avlnode *AVLnode) Insert(value interface{}) *AVLnode {
	if avlnode == nil {
		newNode := &AVLnode{value, nil, nil, 1} //插入节点
		return newNode
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Insert(value)
		avlnode = avlnode.adjust() //自动平衡
	case 1:
		avlnode.Right = avlnode.Right.Insert(value)
		avlnode = avlnode.adjust() //自动平衡
	case 0:
		fmt.Println("数据存在")
	}
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	return avlnode
}

//删除
func (avlnode *AVLnode) Delete(value interface{}) *AVLnode {
	if avlnode == nil {
		return nil
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Delete(value)
	case 1:
		avlnode.Right = avlnode.Right.Delete(value)
	case 0:
		//删除逻辑
		if avlnode.Left != nil && avlnode.Right != nil {
			//左右 都有节点
			avlnode.Data = avlnode.Right.FindMin().Data
			avlnode.Right = avlnode.Right.Delete(avlnode.Data)
		} else if avlnode.Left != nil {
			//左子节点存在,右子节点可能存在也可能不存在
			avlnode = avlnode.Left
		} else {
			avlnode = avlnode.Right
		}
	}
	if avlnode != nil {
		avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
		avlnode = avlnode.adjust()
	}
	return avlnode
}

func (avlnode *AVLnode) Find(data interface{}) *AVLnode {
	var finded *AVLnode = nil
	switch compare(data, avlnode.Data) {
	case -1:
		finded = avlnode.Left.Find(data)
	case 1:
		finded = avlnode.Right.Find(data)
	case 0:
		return avlnode
	default:
		return nil
	}
	return finded
}

func (avlnode *AVLnode) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.height
}

func (avlnode *AVLnode) FindMin() *AVLnode {
	var finded *AVLnode
	//最小值在最左边
	if avlnode.Left != nil {
		finded = avlnode.Left.FindMin() //递归调用
	} else {
		finded = avlnode
	}
	return finded
}

func (avlnode *AVLnode) FindMax() *AVLnode {
	var finded *AVLnode
	//最大值在最右边
	if avlnode.Right != nil {
		finded = avlnode.Right.FindMax() //递归
	} else {
		finded = avlnode
	}
	return finded
}

//抓取数据
func (avlnode *AVLnode) GetData() interface{} {
	return avlnode.Data
}

//设置
func (avlnode *AVLnode) Setdata(data interface{}) {
	avlnode.Data = data
}

func (avlnode *AVLnode) GetLeft() *AVLnode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Left
}

func (avlnode *AVLnode) GetRight() *AVLnode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Right
}
