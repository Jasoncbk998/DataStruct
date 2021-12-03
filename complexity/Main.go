/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/15 9:37 下午
 **/
package main

import "fmt"

func main() {
	//fmt.Println(fib1(5))
	fmt.Println(fib2(10))
}

func fib1(n int) int {
	if n <= 1 && n >= 0 {
		return n
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}

func fib2(n int) int {
	if n <= 1 {
		return n
	} else {
		first := 0
		second := 1
		//var i=0
		for i := 0; i < n-1; i++ {
			sum := first + second
			first, second = second, sum

		}
		return second
	}
}
