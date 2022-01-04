/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/19 5:53 下午
 **/
package main

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func maxSubArray(nums []int) int {
	/**
	假设第一个元素为该数组的最大子序和 称max
	然后将前两个元素的和与max 进行比较 如果大于max,则更新max
	并且,当nums[i]以前的元素和大于nums[i]则会更新nums[i], 对应代码
	*/
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//当nums[i]以前的元素和大于nums[i]则会更新nums[i], 对应代码
		if nums[i]+nums[i-1] > nums[i] {
			//保存第i个元素和i-1元素的和
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func test(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
