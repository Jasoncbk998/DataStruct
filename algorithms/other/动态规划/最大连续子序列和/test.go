/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/8 1:04 下午
 **/
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	array := MaxSubArray_test(nums)
	fmt.Println(array)
}
func MaxSubArray_test(num []int) int {
	if len(num) == 0 {
		return -1
	}
	dp := make([]int, len(num))
	dp[0] = num[0]
	max := dp[0]
	for i := 1; i < len(num); i++ {
		if dp[i-1] <= 0 {
			dp[i] = num[i]
		} else {
			dp[i] = dp[i-1] + num[i]
		}
		max = Max_(dp[i], max)
		fmt.Println("dp[", i, "]:", dp[i])
	}
	return max
}

func Max_(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
