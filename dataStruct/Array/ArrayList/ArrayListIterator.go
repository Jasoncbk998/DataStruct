/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 12:09 上午
 **/
package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool                             //是否有下一个
	Next(password string) (interface{}, error) //下一个
	Remove()                                   //删除
	GetIndex() int                             //获取索引
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}

type ArraylistIterator struct {
	list         *ArrayList // 数组指针
	currentindex int        //当前索引
}

func (it *ArraylistIterator) HasNext() bool {
	return it.currentindex < it.list.TheSize //是否有下一个
}

func (it *ArraylistIterator) Next(password string) (interface{}, error) {
	if password == "111111" {
		if !it.HasNext() {
			return nil, errors.New("没有下一个")
		}
		value, err := it.list.Get(it.currentindex) //抓取当前数据
		it.currentindex++
		return value, err
	}
	return nil, nil
}

func (it *ArraylistIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex) //删除一个元素
}

func (it *ArraylistIterator) GetIndex() int {
	return it.currentindex
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator) //构造迭代器
	it.currentindex = 0
	it.list = list
	return it
}
