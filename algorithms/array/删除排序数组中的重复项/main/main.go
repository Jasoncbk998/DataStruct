/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 3:20 下午
 **/
package main

import (
	test "DataStruct/algorithms/array/删除排序数组中的重复项"
	"fmt"
)

func main() {
	// 要有序
	nums := []int{1, 1, 2, 3, 4, 4}
	duplicates := test.RemoveDuplicates(nums)
	fmt.Println(nums)         // [1 2 3 4 4 4]
	fmt.Println(duplicates)   // 4
	ints := nums[:duplicates] // [1 2 3 4]
	fmt.Println(ints)
}
