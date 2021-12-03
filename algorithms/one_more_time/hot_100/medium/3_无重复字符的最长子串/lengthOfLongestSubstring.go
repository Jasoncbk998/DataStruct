/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 6:37 下午
 **/
package main

/**
Input: "abcabcbb"
Output: 3
滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求。
*/

// 解法二 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	// left right 其实就是两个指针,进行记录次数,通过做差得知最大长度
	result, left, right := 0, 0, -1
	for left < len(s) {
		//出现一次,进行++ right++
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			//发现重复,则-- left++
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

//位图
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
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

//滑动窗口-哈希桶
func lengthOfLongestSubstring(s string) int {
	right, left, res := 0, 0, 0
	indexs := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexs[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexs[s[left]] = left
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

func test(s string) int {
	if len(s) == 0 {
		return 0
	}
	result, left, right := 0, 0, -1
	var freq [256]int
	for left < len(s) {

		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right+1-left)
	}
	return result

}
