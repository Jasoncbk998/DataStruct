/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 1:26 下午
 **/
package main

import (
	_1_searchInsert "DataStruct/algorithms/array/01_searchInsert"
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	insert := _1_searchInsert.SearchInsert1(nums, 7)
	fmt.Println(insert)
	fmt.Println(nums)
}
