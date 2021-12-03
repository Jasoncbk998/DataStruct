/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:24 下午
 **/
package main

import (
	"DataStruct/algorithms/searching"
	"fmt"
)

func main() {
	//一个有序数组
	ints := []int{123, 456, 789, 1234, 12356}
	search := searching.Search(ints, 1234)
	fmt.Println(search)

}
