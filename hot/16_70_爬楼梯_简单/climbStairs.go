/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/22 3:18 下午
 **/
package main

// 简单的 DP，经典的爬楼梯问题。一个楼梯可以由 n-1 和 n-2 的楼梯爬上来。
func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
