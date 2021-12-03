/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 11:06 上午
 **/
package main

import "fmt"

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
*/
func main() {
	s := "abab"
	t := "baba"

	isbool := isAnagram2(s, t)
	fmt.Println(isbool)
}

func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := [26]int{}
	for _, v := range s {
		counts[v-'a']++
	}
	for _, v := range t {
		// s 中的减少次数
		counts[v-'a']--
		// t中进行比对
		if counts[v-'a'] < 0 {
			return false
		}
	}
	return true
}
func isAnagram(s, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}
	schars := []byte(s)
	tchars := []byte(t)
	if len(schars) != len(tchars) {
		return false
	}
	counts := [26]int{}
	for i := 0; i < len(schars); i++ {
		counts[schars[i]-'a']++
	}
	for i := 0; i < len(schars); i++ {
		counts[tchars[i]-'a'] -= 1
		if counts[tchars[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
