/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/12 9:15 下午
 **/
package main

import (
	"fmt"
)

func main() {
	stairs1 := climbStairs1(3)
	fmt.Println(stairs1)
}

// 动态规划
func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		// 这个相当于套公式
		//f(x)=f(x−1)+f(x−2)
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//递归
//通过递归去判断,f(n)=f(n-1)+f(n-2)
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

//递归思路后的优化
func climbStairs3(n int) int {
	if n <= 1 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		second = first + second
		first = second - first
	}
	return second

}
