/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/21 10:52 上午
 **/
package main

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
判断一个数是不是 2 的 n 次方。
这一题最简单的思路是循环，可以通过。但是题目要求不循环就要判断，这就需要用到数论的知识了。这一题和第 326 题是一样的思路。
*/
//如果是2的幂次方,那么%2一定等于0,则降低规模=> n=n/2 然后在%2查看是否等于0
func isPowerOfTwo1(n int) bool {
	for n >= 2 {
		if n%2 == 0 {
			n = n / 2
		} else {
			return false
		}
	}
	return n == 1
}
