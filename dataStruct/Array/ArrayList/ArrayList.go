/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 12:04 上午
 **/
package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newvaal interface{})
	Clear()
	Delete(index int) error
	String() string
	Iterator() Iterator
}
type ArrayList struct {
	DataStore []interface{} //数组存储
	TheSize   int           // 数组大小
}

//new
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.DataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}

//判断是否满
func (list *ArrayList) checkisFull() {
	//判断内存
	if list.TheSize == cap(list.DataStore) {
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) // 双倍内存
		copy(newdataStore, list.DataStore)                                  //进行拷贝数据
		list.DataStore = newdataStore
	}
}

func (list *ArrayList) Size() int {
	return list.TheSize //返回数据大小
}

//追加一个数据
func (list *ArrayList) Append(newval interface{}) {
	list.DataStore = append(list.DataStore, newval) //叠加数据
	list.TheSize++
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.DataStore[index], nil
}

//打印
func (list *ArrayList) String() string {
	return fmt.Sprint(list.DataStore)
}

//设置
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.DataStore[index] = newval //设置
	return nil
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index > list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()                               //是否满,若满则追加
	list.DataStore = list.DataStore[:list.TheSize+1] //插入数据,内存移动一位
	for i := list.TheSize; i > index; i-- {
		//从后往前移动
		list.DataStore[i] = list.DataStore[i-1]
	}
	list.DataStore[index] = newval //插入数据
	list.TheSize++                 //索引追加
	return nil
}

func (list *ArrayList) Clear() {
	list.DataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.TheSize = 0                            //重置大小
}
func (list *ArrayList) Delete(index int) error {
	//删除其实就是跳过需要删除的元素,进行叠加
	list.DataStore = append(list.DataStore[:index], list.DataStore[index+1:]...)
	list.TheSize--
	return nil
}
