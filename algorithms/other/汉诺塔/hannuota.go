/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/4 8:12 下午
 **/
package main

import "fmt"

func main() {
	hanoi(4, "a", "b", "c")
}

//A的盘子挪动到C,每次一个,大得再小的下,且每次挪动大的必须在小的下边
//n 是盘子数量, abc是柱子名称
func hanoi(n int, p1, p2, p3 string) {
	if n == 1 {
		move(n, p1, p3)
		return
	}
	//从p1挪动到p2
	hanoi(n-1, p1, p3, p2)
	move(n, p1, p3)
	hanoi(n-1, p2, p1, p3)
}

func move(n int, from, to string) {
	fmt.Println("将", n, "号盘子从", from, "移动到", to)
}
