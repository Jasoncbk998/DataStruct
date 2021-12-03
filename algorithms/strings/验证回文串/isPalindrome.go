/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 9:09 下午
 **/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	palindrome := isPalindrome1(str)
	fmt.Println(palindrome)
}

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
*/
func isPalindrome(s string) bool {
	var sgood string
	//处理字符串,使其连贯
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}
	//处理后长度
	n := len(sgood)
	//变小写
	sgood = strings.ToLower(sgood)
	//如果是回文串,那么均分后会相等
	for i := 0; i < n/2; i++ {
		//首位和末位进行比对,如果不相等则false
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome1(s string) bool {
	var str string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			str += string(s[i])
		}
	}
	str = strings.ToLower(str)
	n := len(str)

	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}
