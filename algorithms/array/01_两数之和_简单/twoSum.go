/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:41 下午
 **/
package main

/**
01_两数之和_简单
给一个数组,和一个目标值,找到数组内两个值和为目标值得数
*/
//给一个数组,和一个目标值,找到数组内两个值和为目标值得数
func Sum2Number(arr []int, target int) []int {
	maps := map[int]int{}
	for idx, value := range arr {
		if v, ok := maps[target-value]; ok {
			return []int{v, idx}
		}
		maps[value] = idx
	}
	return nil
}
