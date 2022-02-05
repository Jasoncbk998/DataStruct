/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/8 9:27 下午
 **/
package main

/**
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
输入：digits = [1,2,3]
输出：[1,2,4]
*/
func plusOne(digits []int) []int {
	//从后往前取出每一个数字
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
