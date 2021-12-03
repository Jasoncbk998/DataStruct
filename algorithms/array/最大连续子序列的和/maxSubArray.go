/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/21 3:59 下午
 **/
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	array1 := MaxSubArray(nums, 0, len(nums))
	fmt.Println(array1)
}

/**
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

func MaxSubArray2(nums []int) int {
	//取首位元素
	max := nums[0]
	//因为取出首位元素,所以从1开始循环
	for i := 1; i < len(nums); i++ {
		//n位与n-1位作和,与n位比较,看看是否大于,就可以判断是不是局部数组两数之和是不是递增
		if nums[i]+nums[i-1] > nums[i] {
			//赋值给nums[i]
			nums[i] += nums[i-1]
		}
		//更新max
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// MaxSubArray 分治的思想
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return MaxSubArray(nums, 0, len(nums))
}
func MaxSubArray(nums []int, begin, end int) int {
	if end-begin < 2 {
		return nums[begin]
	}
	mid := (begin + end) >> 1
	leftMax := nums[mid-1]
	leftSum := leftMax
	for i := mid - 2; i >= begin; i-- {
		leftSum += nums[i]
		leftMax = max(leftMax, leftSum)
	}
	rightMax := nums[mid]
	rightSum := rightMax
	for i := mid + 1; i < end; i++ {
		rightSum += nums[i]
		rightMax = max(rightMax, rightSum)
	}
	return max(
		leftMax+rightMax,
		max(MaxSubArray(nums, begin, mid), MaxSubArray(nums, mid, end)),
	)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSubArray4(nums []int) int {
	max := nums[0]
	// 模拟窗口滑动
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
