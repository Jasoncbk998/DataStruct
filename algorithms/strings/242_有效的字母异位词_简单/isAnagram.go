/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/29 9:52 上午
 **/
package main

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true
输入: s = "rat", t = "car"
输出: false
*/

func isAnagram(s, t string) bool {
	//把两个字符串的字符都放到码表中
	//然后每出现一次就++,进行比较即可
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
