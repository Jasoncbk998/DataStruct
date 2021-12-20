/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/28 2:55 下午
 **/
package main

import "fmt"

func main() {
	i := reverse(1234)
	fmt.Println(i)
}
func reverse(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10 //取出个位
		if res > 214748364 || (res == 214748364 && temp > 7) {
			return 0
		}
		if res < -214748364 || (res == -214748364 && temp < -8) {
			return 0
		}
		//
		res = res*10 + temp // 很巧妙的,个位提升到十位,十位提升到百位,做到反转
		x /= 10
	}
	return res
}

func rev(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10 //
		res = res*10 + temp
		x /= 10
	}
	return res
}
