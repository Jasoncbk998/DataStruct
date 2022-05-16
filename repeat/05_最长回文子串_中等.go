/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:45 下午
 **/
package main

import "fmt"

func main() {
	s := "bb"
	palindrome_ := longestPalindrome_(s)
	fmt.Println(palindrome_)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	// left right 是检验回文范围左右边界
	// pl pr回文串边界
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		//主要是防止重复字符
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		//边界和回文边界,不断缩小
		//  left-1 >= 0左边界
		// right+1 < len(s) 右边界
		// s[left-1] == s[right+1] 回文判断条件
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

func longestPalindrome__(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
func longestPalindrome_(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// a a a b
		if right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
