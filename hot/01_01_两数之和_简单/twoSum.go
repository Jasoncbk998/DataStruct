/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/19 1:57 下午
 **/
package main

import "fmt"

/**

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
*/
func twoSum(arr []int, target int) []int {
	// arr 2,7,11,15 target = 9
	maps := map[int]int{}
	for index, value := range arr {
		// 9-2=7 maps[7]=1
		if value, ok := maps[target-value]; ok {
			return []int{index, value}
		}
		maps[value] = index
	}
	return nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	sum := twoSum(arr, target)
	fmt.Println(sum)
}
