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
	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))
	//fmt.Println(  [256]bool{})

}

func lengthOfLongestSubstring2(s string) int {
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

// 位图
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	// right 记录无重复, left记录重复
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}
