/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/3 9:04 下午
 **/
package main

/**
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
扫描数组一遍，不断的用 i，j 标记 0 和非 0 的元素，然后相互交换
*/
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
