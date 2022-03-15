/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/16 3:49 下午
 **/
package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum3(ints))
}

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？
请你找出所有和为 0 且不重复的三元组。
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/

// 解法一 最优解，双指针 + 排序
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	start, end, index := 0, 0, 0
	addNum, length := 0, len(nums)
	//每次循环都用index作为第三个数字
	for index = 1; index < length-1; index++ {
		// 头和尾依次与index比较
		start, end = 0, length-1

		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		//以index为中心,开始左右寻找
		// -1, 0, 1, 2, -1, -4
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			// -1, 0, 1, 2, -1, -4
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
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// n1 为负数
		if n1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
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

//-1, 0, 1, 2, -1, -4
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//每次找一个first,然后将这个first与其他所有数字进行匹配测试
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//third从最后一个数字开始
		third := n - 1
		//加和取反,相加才能等于0
		target := -1 * nums[first]
		//从first开始循环,
		for second := first + 1; second < n; second++ {
			//和上次枚举不一样
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//如果指针重合,随着b后续的增加
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}

	}
	return ans
}
