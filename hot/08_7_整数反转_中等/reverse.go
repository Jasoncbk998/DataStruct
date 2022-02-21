/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/21 12:38 下午
 **/
package main

import "fmt"

func reverse(x int) int {
	var result = 0
	// 321
	for x != 0 {
		temp := x % 10
		x /= 10
		result = result*10 + temp
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
}
