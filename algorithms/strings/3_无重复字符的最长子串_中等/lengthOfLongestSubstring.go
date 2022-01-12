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
	fmt.Println(test(s))
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	//right和left交替记录出现次数,然后做减法
	result, left, right := 0, 0, -1
	//从左向右依次比对
	for left < len(s) {
		// right记录出现奇数次,如第一次,第三次
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			//记录一次
			//比如a出现2次,那么这里记录为第一次的记录
			freq[s[right+1]-'a']++
			right++
		} else {
			//left记录偶数次,如第二次出现,第三次
			freq[s[left]-'a']--
			left++
		}
		result = Max(result, right-left+1)
	}
	return result
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func test(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && freq[s[right]-'a'] == 0 {
			freq[s[right]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = Max(result, right-left)
	}
	return result
}
