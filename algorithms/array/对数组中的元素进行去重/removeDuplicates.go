/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/18 10:53 上午
 **/
package main

import "fmt"

/**
Given nums = [1,1,2],
Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

Given nums = [0,0,1,1,1,2,2,3,3,4],
Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
It doesn't matter what values are set beyond the returned length.
*/
func main() {
	//ints := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ints := []int{1, 1, 1, 2, 2}
	duplicates := removeDuplicates(ints)
	fmt.Println(ints[:duplicates])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	//中心思想就是,通过逐位比对,利用两个指针,不断进行替换
	// finder代表了重复的位数
	//比如,1 1 1 2 2 ,finder=3时,发现重复消失,然后进行替换 last+1 = finder(3),可以理解为 0 1 2索引位上有三个重复元素
	// 0, 0, 1, 1
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
