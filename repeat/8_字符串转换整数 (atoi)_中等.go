/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/10 11:43 上午
 **/
package main

import (
	"fmt"
	"math"
)

/*
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
输入：s = "   -42"
输出：-42
解释：
第 1 步："   -42"（读入前导空格，但忽视掉）
            ^
第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -42"（读入 "42"）
               ^
解析得到整数 -42 。
由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42

输入：s = "4193 with words"
输出：4193
解释：
第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
             ^
解析得到整数 4193 。
由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。

输入：s = "-91283472332"
输出：-2147483648
解释：
第 1 步："-91283472332"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："-91283472332"（读入 '-' 字符，所以结果应该是负数）
          ^
第 3 步："-91283472332"（读入 "91283472332"）
                     ^
解析得到整数 -91283472332 。
由于 -91283472332 小于范围 [-231, 231 - 1] 的下界，最终结果被截断为 -231 = -2147483648
*/
func main() {
	s := "-91282"
	atoi := myAtoi(s)
	fmt.Println(atoi)
	fmt.Println(byte('2' - '0'))
}

// 字符串数字转成整数
func myAtoi(s string) int {
	length := len(s)
	runes := []rune(s)
	// 去除空格和其他字符
	index := 0
	// 空格的数量
	for index < length && runes[index] == ' ' {
		index++
	}
	if index == length {
		return 0
	}
	//正负号
	sign := 1
	//依次判断字符
	firstChar := runes[index]
	//正负号判断
	if firstChar == '+' {
		index++
	} else if firstChar == '-' {
		index++
		sign = -1
	}
	//开始拼接
	res := 0
	for index < length {
		currChar := runes[index]
		if currChar > '9' || currChar < '0' {
			break
		}
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && (currChar-'0') > math.MaxInt32%10) {
			return math.MaxInt32
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && (currChar-'0') > -(math.MinInt32%10)) {
			return math.MinInt32
		}
		// 转换思路就是  字节-'0'
		res = res*10 + sign*(int(currChar-'0'))
		index++
	}
	return res
}
