/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/24 3:00 下午
 **/
package main

func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left < right {
		//找到中间值
		mid := left + (right-left)>>1
		//不断从中间值寻找,左右逼近
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
