/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/17 10:26 上午
 **/
package main

import (
	"fmt"
)

/**
给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，使得nums [i] = nums [j]
并且 i 和 j的差的 绝对值 至多为 k。

输入: nums = [1,2,3,1], k = 3
输出: true

输入: nums = [1,0,1,1], k = 1
输出: true
*/
func main() {
	ints := []int{1, 0, 1, 1}
	duplicate := containsNearbyDuplicate(ints, 1)
	fmt.Println(duplicate)
}
func containsNearbyDuplicate1(nums []int, k int) bool {
	m := make(map[int]int)
	for i, val := range nums {
		if _, ok := m[val]; ok == true && i-m[val] <= k {
			return true
		}
		m[val] = i
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := findMax(i-k, 0); j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func findMax(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
