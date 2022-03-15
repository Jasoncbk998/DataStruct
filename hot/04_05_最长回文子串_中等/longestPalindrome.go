/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/19 3:15 下午
 **/
package main

//滑动窗口

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	// 123 333 321
	for left < len(s) {
		//向右寻找重复
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// left向左移动,right像右移动
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		//比对范围进行限定回文边界
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置比较边界
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
