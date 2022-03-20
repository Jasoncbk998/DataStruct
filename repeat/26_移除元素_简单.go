/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/18 9:25 下午
 **/
package main

//nums= 3 2 2 3 val = 3 去除3
func RemoveElement(nums []int, val int) int {
	flag := 0 // flag 当做索引位置
	for _, v := range nums {
		if v != val {
			// 每次遍历,用flag取出数组元素,与目标值进行比对,进行
			nums[flag] = v
			flag++
		}
	}
	return flag
}
