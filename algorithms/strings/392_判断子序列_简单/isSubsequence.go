/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/10 1:14 下午
 **/

package main

//判断s是否为t的子序列
//Input: s = "abc", t = "ahbgdc" true
//Input: s = "axc", t = "ahbgdc"false
func isSubsequence(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
