/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:31 下午
 **/
package main

import "DataStruct/tools"

func lengthOfLongestSubstring_1(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = tools.Max(ans, rk-i+1)
	}
	return ans
}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	// 	result记录无重复个数
	//	left 记录重复次数
	// 	right记录无重复次数
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 从0开始比对
		if right < len(s) && freq[s[right]] == 0 {
			//无重复
			// key 是字符 value是次数
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = Max(result, right-left)
	}
	return result
}

func lengthOfLongestSubstring_s(s string) int {
	result, left, right := 0, 0, 0
	var freq [127]int
	// abcabc
	// 1 1 1
	// 1 1 1
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			right++
			freq[s[right]]++
		} else {
			left++
			freq[s[left]]--
		}
		result = max(result, right-left)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
