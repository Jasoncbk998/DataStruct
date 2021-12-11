/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/15 4:57 下午
 **/
package main

import "fmt"

/**
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
输入：s = "Hello World"
输出：5

输入：s = "   fly me   to   the moon  "
输出：4
*/
func main() {
	s := " e on"
	word := lengthOfLastWord(s)
	fmt.Println(word)
}
func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	//统计出所有字符的个数
	for s[index] == ' ' {
		index--
	}
	//反向遍历,只要遇到第一个空格,那么就找到了最后一个单词的长度
	// 找到字符不为空格 且 ans++
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
