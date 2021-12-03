/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:54 下午
 **/
package main

import (
	"DataStruct/algorithms/array/两数之和"
	"fmt"
)

func main() {
	//arr := []int{11, 22, 33, 44, 66, 11, 33}
	arr := []int{3, 2, 4}
	number := 两数之和.Sum2Number(arr, 6)
	fmt.Println(number)
}
