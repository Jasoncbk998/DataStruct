/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/24 2:28 下午
 **/
package main

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度
用dp
dp[i]代表为第 i 个数字为结尾的最长上升子序列的长度。
换种表述，dp[i] 代表 [0,i] 范围内，选择数字 nums[i] 可以获得的最长上升子序列的长度。
状态转移方程为 dp[i] = max( 1 + dp[j]) ，其中 j < i && nums[j] > nums[i]，取所有满足条件的最大值。
时间复杂度 O(n^2)
*/
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			//上升
			if nums[j-1] < nums[i-1] {
				dp[i] = Max(dp[i], dp[j])
			}
		}
		// 长度
		dp[i] = dp[i] + 1
		res = Max(res, dp[i])
	}
	return res
}
