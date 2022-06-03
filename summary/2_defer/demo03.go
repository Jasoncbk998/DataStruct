/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/3 19:38
 **/
package main

import (
	"errors"
	"fmt"
)

func f1_error() {
	var err error
	defer fmt.Println(err)
	err = errors.New("errors")
	return
}

func f2_error() {
	var err error
	// 闭包,引用了 var err error 变量
	// err 又被赋值了 所以defer打印了erros
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("errors")
	return
}

func f3_error() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err) // 这里传进去的err,其实是err的一个副本,内存地址发生了变化,所以defer打印不是 err=errors.New("errors")
	err = errors.New("errors")
}
func main() {
	f1_error() // nil
	f2_error() // errors
	f3_error() // nil
}
