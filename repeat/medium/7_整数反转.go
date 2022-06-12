/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/25 23:41
 **/
package main

func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		if result > 214748364 || (result == 214748364 && temp > 7) {
			return 0
		}
		//判断是否 小于 最小32位整数
		if result < -214748364 || (result == -214748364 && temp < -8) {
			return 0
		}
		result = result*10 + temp
		x = x / 10
	}
	return result

}
