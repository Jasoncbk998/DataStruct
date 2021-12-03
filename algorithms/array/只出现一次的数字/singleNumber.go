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
	ints := []int{2, 2, 2, 2, 1, 1, 1, 1, 4, 1, 2, 3, 3, 3}
	number := singleNumber_twice(ints)
	fmt.Println(number) // 1
	fmt.Println(0 ^ 2)  // 2
	fmt.Println(2 ^ 0)  // 2
	fmt.Println(2 ^ 2)  // 0
	more := singleNumber_more(ints)
	fmt.Println("more:", more)
}

//位运算
// 如果不是只出现2次重复数字,则可以使用map结构进行
func singleNumber_twice(nums []int) int {
	single := nums[0]
	for i := 1; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}
func singleNumber_more(nums []int) int {
	maps := make(map[int]int)
	//将数组所有元素写进map,key存储数组值,value存储出现次数
	for _, v := range nums {
		if _, ok := maps[v]; ok {
			// 如果重复出现超过2次,则不记录,直接略过
			if maps[v] >= 2 {
				continue
			}
			maps[v] += 1
		} else {
			maps[v] = 1
		}
	}
	//遍历map 发现出现次数为1次 返回
	for k, v := range maps {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber_more4(nums []int) int {
	number := 0
	for _, v := range nums {
		number = number ^ v
	}
	return number

}
