/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/4 12:31 下午
 **/
package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	// result记录无重复个数
	//left 和right分别代表无重复个数&重复的个数
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
