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
//滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left)
	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcbb"
	fmt.Println(lengthOfLongestSubstring2(s))
}

func lengthOfLongestSubstring2(s string) int {
	m := map[byte]int{}
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		//边界是 right+1<n 并且 map中不存在该元素
		//一直添加到遇到重复的字符
		for right+1 < n && m[s[right+1]] == 0 {
			bytes := s[right+1]
			m[bytes]++
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}
