/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 12:06 下午
 **/
package main

import (
	"DataStruct/algorithms/array/02_removeElement"
	"fmt"
)

func main() {
	nums := []int{3, 4, 3, 5, 3, 6}
	element := _2_removeElement.RemoveElement1(nums, 3)
	fmt.Println(element)        // 3
	fmt.Println(nums[:element]) // 4 5 6
	fmt.Println(nums)           // [4 5 6 5 3 6]
}
