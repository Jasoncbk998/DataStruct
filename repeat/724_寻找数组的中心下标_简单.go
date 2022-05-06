/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/2 16:17
 **/
package main

/**
输入：nums = [1, 7, 3, 6, 5, 6]
输出：3
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等
假设,左侧元素和是sum,现在遍历到i,数组和是total
那么左右相等,则
sum=total-num[i]-sum
2sum+num[i]=total,遍历这个公式就行了,简单题
*/
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
