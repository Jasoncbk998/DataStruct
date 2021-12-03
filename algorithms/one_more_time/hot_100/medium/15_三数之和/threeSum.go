/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/27 12:31 下午
 **/
package main

import "sort"

/**
三个数相加为0
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
// 最好理解的
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//left是i的后一个值,right是最后一个值
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

//双指针+排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start--
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
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
