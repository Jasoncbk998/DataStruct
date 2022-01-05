/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 9:39 下午
 **/
package main

/**
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
输入: [1,2,3,1]
输出: true
*/
//利用map 一次for
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		if m[v] == 1 {
			return true
		} else {
			m[v] = 1
		}
	}
	return false
}
