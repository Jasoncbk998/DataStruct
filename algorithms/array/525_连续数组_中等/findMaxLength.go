/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/23 4:46 下午
 **/
package main

/**
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
*/
func findMaxLength(nums []int) int {
	dict := map[int]int{}
	dict[0] = -1
	count, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		} else {
			count++
		}
		if idx, ok := dict[count]; ok {
			res = max(res, i-idx)
		} else {
			dict[count] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
