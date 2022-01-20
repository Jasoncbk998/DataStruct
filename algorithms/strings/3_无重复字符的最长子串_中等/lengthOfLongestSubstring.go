/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/17 7:01 下午
 **/
package main

import "fmt"

/**
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/
func main() {
	s := "abcaaaa"
	fmt.Println(lengthOfLongestSubstring(s))
}

//滑动窗口-哈希桶
// aabcddac
func lengthOfLongestSubstring(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
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
