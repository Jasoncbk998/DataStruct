/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/17 10:26 上午
 **/
package main

/**
给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，使得nums [i] = nums [j]
并且 i 和 j的差的 绝对值 至多为 k。

输入: nums = [1,2,3,1], k = 3
输出: true

输入: nums = [1,0,1,1], k = 1
输出: true
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, val := range nums {
		if _, ok := m[val]; ok == true && i-m[val] <= k {
			return true
		}
		m[val] = i
	}
	return false
}
