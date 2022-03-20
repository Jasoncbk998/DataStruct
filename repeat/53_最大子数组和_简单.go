/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:48 下午
 **/
package main

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func maxSubArray(nums []int) int {
	//假定第一个元素是最大和
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//如果当前位+前位>当前位,说明可能发现较大值,将该较大值保存至当前位
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		// 当前位和max进行比较更新
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
