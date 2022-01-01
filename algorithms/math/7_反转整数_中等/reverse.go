/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/14 9:32 下午
 **/
package main

func reverse(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10 // 取末尾数字
		if res > 214748364 || (res == 214748364 && temp > 7) {
			return 0
		}
		//判断是否 小于 最小32位整数
		if res < -214748364 || (res == -214748364 && temp < -8) {
			return 0
		}
		// res默认值是0
		//相当于每一次循环就*10一次 然后加上temp,而temp取的是x的末尾数字
		res = res*10 + temp
		// 如果x是4位数字,第一次除10 则变为3位数
		x /= 10
	}
	return res
}
