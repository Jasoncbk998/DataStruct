/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/28 10:38 上午
 **/
package main

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	idx := 0 // 为0元素
	// 5,0,0,1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx += 1
		}
	}
}
