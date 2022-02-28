/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/26 2:48 下午
 **/
package main

/**
输入：x = 4
输出：2
*/

func mySqrt(x int) int {
	l, r := 0, x
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
