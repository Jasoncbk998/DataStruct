/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 7:09 下午
 **/
package main

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

func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, end := 0, 0

	for i := 0; i < n; {
		l, r := i, i
		//如果字符串相同则分别冲 一个和后一个开始回文
		for r < n-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if end < r-l {
			start = l
			end = r - l
		}
	}
	return s[start : start+end+1]

}
