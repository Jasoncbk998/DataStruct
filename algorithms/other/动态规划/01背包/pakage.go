/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/8 9:02 下午
 **/
package main

import (
	"fmt"
	"math"
)

/**
有n件物品和一个最大承重为W的背包,每件物品的重量是w,价值是v
保证总重量不超过w的前提下,将哪些物品放进背包,使总价值最大
每个物品只有1件,所以叫01背包
*/
func main() {
	values := []int{6, 3, 5, 4, 6}
	weights := []int{2, 2, 6, 5, 4}
	capacity := 10
	value1 := maxValue2(values, weights, capacity)
	fmt.Println(value1)

}

//恰好凑到capacity容量
//如果返回-1，代表没法刚好凑到capacity这个容量
func maxValueExactly(values, weights []int, capacity int) int {
	if len(values) == 0 || len(weights) == 0 {
		return 0
	}
	if len(values) != len(weights) || capacity <= 0 {
		return 0
	}
	dp := make([]int, capacity+1)

	for j := 1; j <= capacity; j++ {
		dp[j] = math.MinInt64
	}

	if dp[capacity] < 0 {
		return -1
	} else {
		return dp[capacity]
	}

}

func maxValue2(values, weights []int, capacity int) int {
	if len(values) == 0 || len(weights) == 0 {
		return 0
	}
	if len(values) != len(weights) || capacity <= 0 {
		return 0
	}
	dp := make([]int, capacity+1)
	for i := 0; i <= len(values); i++ {
		for j := capacity; j >= 1; j-- {
			if j < weights[i-1] {
				continue
			}
			dp[j] = Max(dp[j], values[i-1]+dp[j-weights[i-1]])
		}
	}
	return dp[capacity]

}
func maxValue1(values, weights []int, capacity int) int {
	if len(values) == 0 || len(weights) == 0 {
		return 0
	}
	if len(values) != len(weights) || capacity <= 0 {
		return 0
	}

	dp := make([][]int, len(values)+1)
	for i := 0; i < len(values)+1; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(values); i++ {
		for j := 1; j <= capacity; j++ {
			if j < weights[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = Max(dp[i-1][j],
					values[i-1]+dp[i-1][j-weights[i-1]])
			}
		}
	}
	return dp[len(values)][capacity]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
