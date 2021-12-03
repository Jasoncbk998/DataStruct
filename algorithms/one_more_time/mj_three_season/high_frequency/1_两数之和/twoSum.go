/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 3:45 下午
 **/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
func twoSum(arr []int, target int) []int {
	maps := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		temp := target - arr[i]
		if value, ok := maps[temp]; ok {
			return []int{value, i}
		}
		maps[arr[i]] = i
	}
	return nil
}
