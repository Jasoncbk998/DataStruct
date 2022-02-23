/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/23 11:34 上午
 **/
package main

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。问总共有多少条不同的路径？

只能向下,或者向右即(i-1,j),(i,j-1)
动态转移方程: f(i,j)=f(i−1,j)+f(i,j−1)
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	//相当于画一个平面图,长和宽分别是m和n
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 //一直向右走
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 一直像下走
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// f(i,j)=f(i−1,j)+f(i,j−1)
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
