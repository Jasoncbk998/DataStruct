/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/8 9:27 下午
 **/
package main

import "fmt"

/**
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
*/
func plusOne1(digits []int) []int {
	//给定一个数组
	// 从后往前进行遍历,缝9进1 所以从后往前循环
	for i := len(digits) - 1; i >= 0; i-- {
		//末位是9则直接置为0
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			// 不是9,则+1
			digits[i] += 1
			return digits
		}
	}
	//处理个位数且个位数是9
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
func main() {
	//arr := []int{3, 9}
	arr := []int{1, 3, 4, 9}
	one := plusOne2(arr)
	fmt.Println(one)
}

//最优解
func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		//逢9进1,然后
		if digits[i] != 10 {
			return digits
		}
		// carry
		digits[i] = 0
	}
	// all carry
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}

// 1 2 9
// 1 3 0
func plusOne3(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i]++
		if nums[i] != 10 {
			return nums
		}
		nums[i] = 0
	}
	nums[0] = 1
	nums = append(nums, 0)
	return nums
}
