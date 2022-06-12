/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:31 下午
 **/
package main

import (
	"fmt"
)

func main() {
	str := "abcd"
	substring_1 := lengthOfLongestSubstring_1(str)
	fmt.Println(substring_1)
}
func lengthOfLongestSubstring_1(s string) int {
	// map的key是字符,valye是次数
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	// 以s[i]开头的最长无重复子串
	for i := 0; i < n; i++ {
		if i != 0 {
			kyes := s[i-1]
			delete(m, kyes)
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
