/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 9:39 下午
 **/
package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	duplicate := containsDuplicate3(ints)
	fmt.Println(duplicate)
}

/**
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
输入: [1,2,3,1]
输出: true
*/
func containsDuplicate(nums []int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			//相同元素 ^运算结果=0 ,通过i!=j 控制不要比较同一个元素
			if nums[i]^nums[j] == 0 {
				return true
			}
		}
	}
	return false
}
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

//很慢
func containsDuplicate3(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

//巧妙
func containsDuplicate4(nums []int) bool {
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
