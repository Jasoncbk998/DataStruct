/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 1:04 下午
 **/
package main

/**
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

}

//动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i]
		} else {
			dp[i] = nums[i]
		}
		res = Max(res, dp[i])
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray_test(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return 0

}
