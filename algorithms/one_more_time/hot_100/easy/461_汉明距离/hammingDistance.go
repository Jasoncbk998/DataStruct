/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 11:45 上午
 **/
package main

/**
求 2 个数的海明距离。海明距离的定义是两个数二进制位不同的总个数。
这一题利用的位操作的是 X &= (X - 1) 不断的清除最低位的 1 。
先将这两个数异或，异或以后清除低位的 1 就是最终答案。
*/
func hammingDistance(x int, y int) int {
	var ans int
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			ans += 1
		}
		x /= 2
		y /= 2
	}
	return ans
}
