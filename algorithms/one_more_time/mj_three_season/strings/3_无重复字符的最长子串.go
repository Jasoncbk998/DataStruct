/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 1:20 下午
 **/
package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
func lengthOfLongestSubstring(s string) int {
	maps := make(map[byte]int, len(s))
	left, right, res := 0, 0, 0
	for left < len(s) {
		if idx, ok := maps[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		maps[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
