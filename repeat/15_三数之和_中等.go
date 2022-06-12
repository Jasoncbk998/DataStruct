/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/19 4:00 下午
 **/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, -3, 3, -1, -4, 4, 5}
	sum := threeSum(nums)
	fmt.Println(sum)
}

//	输入：nums = [-1,0,1,2,-1,-4]
//	输出：[[-1,-1,2],[-1,0,1]]
// 	双指针+排序
func threeSum(nums []int) [][]int {
	//nums = [-1,0,1,2,-1,-4]
	sort.Ints(nums) //有序
	// -4,-1,-1,0,1,2
	result, length := make([][]int, 0), len(nums)
	addNum := 0
	start, end := 0, 0
	index := 0
	// index从序号1开始
	// 每次分别与start,end 进行求和比较
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		// 重复数字
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		//头,中,尾 三数进行求和比较
		for start < index && end > index {
			//相同则下一位
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			// 求和
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func ____threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, lengh := make([][]int, 0), len(nums)
	addNum := 0
	start, end := 0, 0
	index := 0
	for index = 1; index < lengh-1; index++ {
		start, end = 0, lengh-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < lengh-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[end], nums[index]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}

		}
	}

	return result
}
