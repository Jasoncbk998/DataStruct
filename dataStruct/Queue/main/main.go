/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 8:13 下午
 **/
package main

import (
	"DataStruct/dataStruct/Queue"
	"fmt"
	"io/ioutil"
)

func main1() {
	myq := Queue.NewQueue()
	myq.EnQueue("1")
	fmt.Println("s", myq.DeQueue())
	myq.EnQueue("2")
	fmt.Println("s", myq.DeQueue())
	myq.EnQueue("3")
	fmt.Println("s", myq.DeQueue())
}

func main() {
	files := []string{}
	path := "dataStruct/LinkedLIst"
	queue := Queue.NewQueue()
	queue.EnQueue(path)
	for {
		if !queue.IsEmpty() {
			//出队列
			path = queue.DeQueue().(string)
		}
		if path == "" {
			break
		}
		read, err := ioutil.ReadDir(path)
		if err != nil {
			break
		}
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name()
				//进队列
				queue.EnQueue(fulldir)

			} else {
				fullname := path + "/" + fi.Name()
				files = append(files, fullname)
			}
		}
		path = ""
	}
	for _, v := range files {
		fmt.Println(v)
	}

}
