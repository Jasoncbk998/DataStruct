/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/22 3:56 下午
 **/
package main

import "fmt"

func main() {
	arr := make([]int, 5, 10)
	i := arr[0]
	i1 := arr[1]
	i2 := arr[2]
	fmt.Println(i)  // 0
	fmt.Println(i1) // 0
	fmt.Println(i2) // 0

	//arr2 := new([]int) // 不会赋默认值,只是new一个对象,没有分配内存
	//a1 := (*arr2)[0]  // 会panic
	//fmt.Println(a1)
}
