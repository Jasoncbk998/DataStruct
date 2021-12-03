/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 7:07 下午
 **/
package main

import "fmt"

func main() {
	fmt.Println(reverse7(123))
}
func reverse7(x int) int {
	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	if temp > 1<<31-1 || temp < -(1<<31) {
		return 0
	}
	return temp

}
