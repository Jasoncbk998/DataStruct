/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 10:26 上午
 **/
package main

import (
	merge_sort "DataStruct/algorithms/sorting/merge-sort"
	"fmt"
)

func main() {
	array := []int{5, 4, 3, 2, 1}
	merge_sort.Sort(array)
	for k, v := range array {
		fmt.Println(k, v)
	}
}
