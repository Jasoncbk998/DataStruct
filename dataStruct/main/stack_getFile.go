/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 2:08 下午
 **/
package main

import (
	"DataStruct/dataStruct/Stack/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

/**
stack 栈
是深度遍历
比如a里边有很多层文件夹,他会直接从一个路径下挖到最深

queue
而队列则不是这样,队列是把同一级所有文件夹都遍历出来 再去遍历下一级目录
这个就是深度和广度遍历的区别


递归就是深度遍历   可以用栈来模拟


*/
func main() {
	path := "/Users/liuxu/GolandProjects/DataStruct"
	files := []string{}
	mystack := StackArray.NewStack()
	//把第一个路径压栈
	mystack.Push(path)
	for !mystack.IsEmpty() {
		//弹出
		path := mystack.Pop().(string)
		files = append(files, path)
		//读取路径
		read, _ := ioutil.ReadDir(path)
		//继续遍历
		for _, fi := range read {
			//如果是文件夹,继续压栈
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name()
				mystack.Push(fulldir)
			} else {
				//如果是文件直接append到数组中
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}
	for _, v := range files {
		fmt.Println(v)
	}
}

// 递归文件夹至目录
func main_digui() {
	files := []string{}
	all, _ := GetAll("/Users/liuxu/GolandProjects", files)
	for _, v := range all {
		fmt.Println(v)
	}
}

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件不可读取")
	}
	for _, fi := range read {
		//是文件夹,进入递归
		if fi.IsDir() {
			fulldir := path + "/" + fi.Name()
			files = append(files, fulldir)
			files, _ = GetAll(fulldir, files) //开始递归
		} else {
			//不是文件夹,追加到数组内
			fulldis := path + "/" + fi.Name()
			files = append(files, fulldis)
		}
	}
	return files, nil
}
