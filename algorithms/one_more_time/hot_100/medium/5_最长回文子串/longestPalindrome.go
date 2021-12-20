/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 7:09 下午
 **/
package main

import "fmt"

/**
Input: s = "babad"
Output: "bab"
*/
// DP，时间复杂度 O(n^2)，空间复杂度 O(n^2)
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && ((res == "") || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res

}

func main() {
	lines := longestPalindrome1("eeeabccbaccc")
	fmt.Println(lines)
}

// 滑动窗口
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	//left right 是寻找的左右边界,确定的最长回文子串的边界是pl和pr
	//right=-1 是一个简单的技巧,-1开始,可以通过+1操作去找到索引位置是0的元素
	left, right, pl, pr := 0, -1, 0, 0
	//从左边界left开始循环
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）,确定边界,left到right
		//从left开始寻找,找到right的位置 (如果有相同字母的情况)
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 在边界中找到回文的子串边界
		//从左边界开始,left--与rigjt++进行比较,
		//循环条件就是left-1>=0 right+1<len(s) 不能越界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		//更新回文串子串边界
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心,更新边界
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
