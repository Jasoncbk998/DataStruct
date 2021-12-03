/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/30 10:23 上午
 **/
package main

import "fmt"

func main() {
	nums := []int{5, 0, 0, 2, 3}
	moveZeroes(nums)
	fmt.Println(nums)

}

//双指针
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0 //代表0的位置
	// 0 0 4 2 1
	//遍历nums中的每一个元素,
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
}
