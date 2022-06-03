/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/25 22:12
 **/
package main

/**
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	// right:非重复
	// left:重复
	result, left, right := 0, 0, 0
	var freq [127]int
	for left < len(s) {
		// 非重复
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			//重复
			freq[s[left]]--
			left++
		}
		// 求max
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
