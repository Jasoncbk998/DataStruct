/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/9 11:51 下午
 **/
package main

import (
	"fmt"
	"math"
)

/**
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
输入：x = 4
输出：2
*/
func main() {
	sqrt := mySqrt(4)
	fmt.Println(sqrt)

}

//二分查找
func mySqrt(x int) int {
	//确定左右边界
	l, r := 0, x
	ans := -1 // 设置默认值
	//循环查找
	for l <= r {
		//寻找中间点
		mid := l + (r-l)/2
		// 不断地用中间值*中间值 判断是否<= x,右临界点
		if mid*mid <= x {
			//如果小于,则将本次的mid赋值给ans,并且右移l,不断夹逼r
			ans = mid
			l = mid + 1
		} else {
			//反之,左移r,夹逼l
			r = mid - 1
		}
	}
	return ans
}

// 牛顿迭代
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
