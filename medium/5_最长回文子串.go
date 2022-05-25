/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/25 23:23
 **/
package main

func longestPalindrome(s string) string {
	pl, pr, left, right := 0, 0, 0, -1
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[right+1] == s[left-1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pr = right
			pl = left
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
