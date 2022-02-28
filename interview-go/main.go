/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/24 8:46 下午
 **/
package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	a := "aaa"
	fmt.Println("a:", &a)

	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Println(ssh)
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	// b和a的值是一样的
	fmt.Println(b)
	println(strings.EqualFold(a, string(b))) // true
	pointer := unsafe.Pointer(&b)
	fmt.Println(pointer)

}
