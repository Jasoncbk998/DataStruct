/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 3:08 下午
 **/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//层级递归
	path := "/Users/liuxu/"
	GetAll_NB(path, 1)
}

func GetAll_NB(path string, level int) {
	levelstr := ""
	if level == 1 {
		levelstr = "---"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "+"
	}
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "/" + fi.Name()
			fmt.Println(levelstr + fulldir)
			GetAll_NB(fulldir, level+1)
		} else {
			fulldir := path + "\\" + fi.Name() //构造新的路径
			fmt.Println(levelstr + fulldir)
		}
	}
}
