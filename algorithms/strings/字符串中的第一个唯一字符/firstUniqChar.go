/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/28 9:02 下午
 **/
package main

import "fmt"

/**
给定一个字符串
找到它的第一个不重复的字符，并返回它的索引。
如果不存在，则返回 -1。
s = "leetcode"
返回 0
你可以假定该字符串只包含小写字母。
s = "loveleetcode"
返回 2
*/

func main() {
	s := "abc"
	char := firstUniqChar(s)
	fmt.Println(char)
}
func firstUniqChar(s string) int {
	// 模拟26个英文字母,相当于码表
	cnt := [26]int{}
	for _, ch := range s {
		// ch-'a' a是0 b是1 c是2 以此类推,比如 cnt[1]=2 那么b出现了2次
		cnt[ch-'a']++
	}
	for i, ch := range s {
		//找出出现一次的字母
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
