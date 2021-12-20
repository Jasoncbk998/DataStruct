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

//   滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	//码表
	var freq [256]int
	// left right 其实就是两个指针,进行记录次数,通过做差得知最大长度
	result, left, right := 0, 0, -1
	//从left开始寻找,依次比较
	//没发现重复就,扩大右边界,发现重复就缩小左边界
	for left < len(s) {
		//right+1 依次校验是否出现过
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			//发现重复,则缩小左边界,进行left++
			freq[s[left]-'a']--
			left++
		}
		//每次元素都进行一次判断边界长度
		result = max(result, right-left+1)
	}
	//需要一次遍历
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
