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
	// 创建两个码表
	// a的索引就是0,b是1 以此类推,c1[1]=1 就是a出现一次
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
