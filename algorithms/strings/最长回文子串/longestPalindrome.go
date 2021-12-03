/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 1:03 下午
 **/
package main

import "fmt"

/**
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func main() {
	s := "babad"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

//func longestPalindrome2(s string) string {
//		length:=len(s)
//	if length<2{
//		return s
//	}
//	maxLen:=1
//	begin:=0
//
//	dp := make([][]int, length, length)
//
//	for i := 0; i < len(char1)+1; i++ {
//		dp[i] = make([]int, len(char2)+1)
//	}
//}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]

}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
