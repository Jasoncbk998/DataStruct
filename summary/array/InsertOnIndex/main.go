/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 1:38 下午
 **/
package main

import "fmt"

func main() {
	// 在切片中的指定index加入元素
	// 在index=2 处加入3
	sliceHaiCoder := []int{1, 2, 4, 5, 6, 7, 8}
	sliceHaiCoder = append(sliceHaiCoder[:2], append([]int{3}, sliceHaiCoder[2:]...)...)
	fmt.Println(sliceHaiCoder)
}
