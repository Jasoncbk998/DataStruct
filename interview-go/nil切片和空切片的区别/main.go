/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/23 8:53 下午
 **/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	// s1 pointer:{Data:0 Len:0 Cap:0}, s2 pointer:{Data:824634216112 Len:0 Cap:0}, s4 pointer:{Data:824634216112 Len:0 Cap:0},
	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
