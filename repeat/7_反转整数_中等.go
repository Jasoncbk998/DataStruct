/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/6 4:41 下午
 **/
package main

func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		x /= 10
		result = result*10 + temp
	}
	return result
}

func reverse_(x int) int {
	result := 0
	// 123
	for x != 0 {
		temp := x % 10
		x /= 10
		result = result*10 + temp
	}
	return result
}
