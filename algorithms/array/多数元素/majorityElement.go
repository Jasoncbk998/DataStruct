/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 10:18 上午
 **/
package main

import "fmt"

func main() {
	nums := []int{3, 3, 3, 1, 2, 2, 2, 2, 2, 2, 2, 2}
	element := majorityElement(nums)
	fmt.Println(element)
}

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
输入：[2,2,1,1,1,2,2]
输出：2
*/
func majorityElement(nums []int) int {
	maps := make(map[int]int)
	for _, v := range nums {
		_, ok := maps[v]
		if ok {
			maps[v] = maps[v] + 1
		} else {
			maps[v] = 1
		}
	}
	count := 0
	temp := 0
	for k, v := range maps {
		if v >= count {
			temp = k
			count = v
		}
	}
	return temp
}

//逐个元素进行比较
func majorityElement1(nums []int) int {
	isCddEmpty := true
	candidate := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if isCddEmpty {
			candidate = nums[i]
			count = 1
			isCddEmpty = false
		} else {
			if candidate == nums[i] {
				count += 1
			} else {
				count -= 1
				if count == 0 {
					isCddEmpty = true
				}
			}
		}
	}
	return candidate
}
