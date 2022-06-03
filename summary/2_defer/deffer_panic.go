/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/3 19:54
 **/
package main

import "fmt"

func main() {
	defer t()
	panic(any(404))
	fmt.Println("00") // 不会执行
}
func t() {
	if e := recover(); e != any(nil) {
		fmt.Println("recover")
		return
	}
}
