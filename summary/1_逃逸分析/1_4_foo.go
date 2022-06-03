/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/3 18:04
 **/
package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func main() {
	x := foo()
	fmt.Println(x)
}
