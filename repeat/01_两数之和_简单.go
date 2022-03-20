/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/20 1:54 下午
 **/
package main

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
