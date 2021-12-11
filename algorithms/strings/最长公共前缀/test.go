/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/12 10:12 上午
 **/
package 最长公共前缀

func LongPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := len(strs)
	prefix := strs[0]

	for i := 1; i < length; i++ {
		prefix = findPrefix(prefix, strs[i])
		if (len(prefix)) == 0 {
			break
		}
	}
	return prefix

}

func findPrefix(str1 string, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index <= length && str1[:index] == str2[:index] {
		index++
	}
	return str1[:index]

}
