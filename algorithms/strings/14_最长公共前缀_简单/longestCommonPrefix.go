/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 1:48 下午
 **/
package main

/**
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		//更新前缀
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	//前包后不包
	return str1[:index]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
