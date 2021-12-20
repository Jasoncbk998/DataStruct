/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/19 6:55 下午
 **/
package main

func main() {
	println(mySqrt(36))
}
func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left < right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
