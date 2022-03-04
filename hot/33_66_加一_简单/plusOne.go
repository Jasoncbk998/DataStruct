/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/3 8:58 下午
 **/
package main

/**
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
*/

func plusOne(digits []int) []int {
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
