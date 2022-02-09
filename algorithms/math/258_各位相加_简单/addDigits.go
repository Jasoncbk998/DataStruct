/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/9 9:05 下午
 **/
package main

import "fmt"

/**
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数
定一个非负整数，反复加各个位上的数，直到结果为一位数为止，最后输出这一位数。
简单题。按照题意循环累加即可。
*/
func main() {
	digits := addDigits(99)
	fmt.Println(digits)

}
func addDigits(num int) int {
	for num > 9 {
		cur := 0
		for num != 0 {
			//取出来个位
			cur += num % 10
			//降位数  百位->十位 十位->个位
			num /= 10
		}
		num = cur
	}
	return num
}
