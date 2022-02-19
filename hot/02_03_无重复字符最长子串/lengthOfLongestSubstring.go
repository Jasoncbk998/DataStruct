/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/19 2:05 下午
 **/
package main

import "fmt"

/**
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func main() {
	s := "abba"
	substring := lengthOfLongestSubstring2(s)
	fmt.Println(s[:substring])
}

// 另一种滑动窗口
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	// right	代表非重复次数
	// left 	代表重复次数
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			//没重复
			freq[s[right]]++
			right++
		} else {
			//重复
			freq[s[left]]--
			left++
		}
		result = max(result, right-left)
	}
	return result
}

//滑动窗口的经典解法
/**
以 \texttt{(a)bcabcbb}(a)bcabcbb 开始的最长字符串为 \texttt{(abc)abcbb}(abc)abcbb；
以 \texttt{a(b)cabcbb}a(b)cabcbb 开始的最长字符串为 \texttt{a(bca)bcbb}a(bca)bcbb；
以 \texttt{ab(c)abcbb}ab(c)abcbb 开始的最长字符串为 \texttt{ab(cab)cbb}ab(cab)cbb；
以 \texttt{abc(a)bcbb}abc(a)bcbb 开始的最长字符串为 \texttt{abc(abc)bb}abc(abc)bb；
以 \texttt{abca(b)cbb}abca(b)cbb 开始的最长字符串为 \texttt{abca(bc)bb}abca(bc)bb；
以 \texttt{abcab(c)bb}abcab(c)bb 开始的最长字符串为 \texttt{abcab(cb)b}abcab(cb)b；
以 \texttt{abcabc(b)b}abcabc(b)b 开始的最长字符串为 \texttt{abcabc(b)b}abcabc(b)b；
以 \texttt{abcabcb(b)}abcabcb(b) 开始的最长字符串为 \texttt{abcabcb(b)}abcabcb(b)。
*/
func lengthOfLongestSubstring1(s string) int {
	m := map[byte]int{}
	n := len(s)
	left, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for left+1 < n && m[s[left+1]] == 0 {
			m[s[left+1]]++
			left++
		}
		ans = max(ans, left+1-i)
	}
	return ans
}

// 这个解法不是很理解,有一点难懂
func lengthOfLongestSubstring(s string) int {
	maps := make(map[byte]int, len(s))
	right, left, res := 0, 0, 0
	// abcabc
	for left < len(s) {
		//
		if idx, ok := maps[s[left]]; ok && idx >= right {
			//出现重复
			right = idx + 1
		}
		maps[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
