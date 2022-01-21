/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/21 1:02 下午
 **/
package main

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度
用dp
*/
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
