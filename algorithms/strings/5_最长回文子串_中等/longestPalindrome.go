/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 1:03 下午
 **/
package main

import "fmt"

/**
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func main() {
	s := "babad"
	palindrome := longestPalindrome1(s)
	fmt.Println(palindrome)
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
