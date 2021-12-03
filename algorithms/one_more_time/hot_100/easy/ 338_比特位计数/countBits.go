/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 11:55 上午
 **/
/**
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
给出一个数，要求计算出 0 ≤ i ≤ num 中每个数的二进制位 1 的个数。

这一题就是利用二进制位运算的经典题。
X&1==1or==0，可以用 X&1 判断奇偶性，X&1>0 即奇数。
X = X & (X-1) 清零最低位的1
X & -X => 得到最低位的1
X&~X=>0
*/
package main

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] += bits[i&(i-1)] + 1
	}
	return bits
}
