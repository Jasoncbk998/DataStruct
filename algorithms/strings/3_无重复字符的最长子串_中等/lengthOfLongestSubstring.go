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

//双指针配合码表,进行夹逼
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		// 因为right 从-1开始
		//因为byte索引是0开始,所以方便计算right=-1
		// 比如aaaa 第一次进入freq++ 为1,right++ right=1
		//第二次,ferq-- 为0 ,left++ left=1
		//第三次,freq++ 为1
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			//记录一次
			//比如a出现2次,那么这里记录为第一次的记录
			freq[s[right+1]-'a']++
			right++
		} else {
			//记录第二次 or 第n次出现的次数
			//每进入一次,数组就--
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
