/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 4:09 下午
 **/
package main

import (
	"DataStruct/dataStruct/Queue"
	"fmt"
	"io/ioutil"
)

/**
队列   广度遍历   不是深度遍历
递归文件夹
*/

func main() {
	path := "/Users/liuxu/GolandProjects/DataStruct/algorithms/array/01_searchInsert"
	files := []string{}

	queue := Queue.NewQueue()
	queue.EnQueue(path)
	for {
		path := queue.DeQueue()
		if path == nil {
			break
		}
		//fmt.Println("get", path)
		read, _ := ioutil.ReadDir(path.(string))
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string) + "/" + fi.Name()
				//fmt.Println("Dir", fulldir)
				queue.EnQueue(fulldir) //继续进队列
			} else {
				fulldir := path.(string) + "/" + fi.Name()
				files = append(files, fulldir)
				//fmt.Println("File", fulldir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}
