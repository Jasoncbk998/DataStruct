/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/14 9:50 下午
 **/
package main

import "fmt"

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
输入: [4,1,2,1,2]
输出: 4
*/
func main() {
	//ints := []int{2, 2, 2, 2, 1, 1, 1, 1, 4, 1, 2, 3, 3, 3}
	fmt.Println(0 ^ 2) // 2
	fmt.Println(2 ^ 0) // 2
	fmt.Println(2 ^ 2) // 0
}

func singleNumber(nums []int) int {
	maps := make(map[int]int)
	//遍历数组,放到map中,重复的数字大于1
	for _, v := range nums {
		if _, ok := maps[v]; ok {
			maps[v] += 1
		} else {
			maps[v] = 1
		}
	}
	for k, v := range maps {
		if v == 1 {
			return k
		}
	}
	return 0
}
