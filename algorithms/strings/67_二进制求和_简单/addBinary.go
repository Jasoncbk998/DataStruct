/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/6 1:09 下午
 **/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "11"
	b := "1"
	binary := addBinary(a, b)
	fmt.Println(binary)
}

/**
输入: a = "11", b = "1"
输出: "100"
*/
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
			fmt.Println(carry)
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
