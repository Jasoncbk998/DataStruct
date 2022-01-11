/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/28 10:38 上午
 **/
package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func main() {
	ints := []int{0, 0, 1, 0, 2}
	moveZeroes(ints)
	fmt.Println(ints)
}
func moveZeroes(nums []int) {
	//right代表非零
	//left代表为零
	left, right, n := 0, 0, len(nums)
	for right < n {
		//非零
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			//因为对0进行换位,所以为零元素进行++
			left++
		}
		//right为0,则++
		right++
	}
}
