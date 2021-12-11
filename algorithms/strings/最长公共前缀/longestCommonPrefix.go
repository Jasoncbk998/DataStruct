/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 1:48 下午
 **/
package 最长公共前缀

/**
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//假设数组第一个值是最小前缀
	prefix := strs[0]
	//判断数组长度
	count := len(strs)
	//开始循环,交替比较字符串
	//因为默认取数组第一个元素,所以从第二个元素开始比较
	//不断的循环比较
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
	//进行比较,找到最小长度
	length := min(len(str1), len(str2))
	index := 0
	//找到前缀的最大索引位置
	for index < length && str1[index] == str2[index] {
		index++
	}
	//前包后不包
	//找到比较元素的共有前缀,返回
	return str1[:index]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
