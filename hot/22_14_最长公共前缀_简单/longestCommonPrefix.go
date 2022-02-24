/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/24 11:52 上午
 **/
package main

/**
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
/**
纵向扫描
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//str:=[]string{"flower","flow","flight"}
	//u := str[0][0]
	//fmt.Println(string(u))// f

	// strs = ["flower","flow","flight"]
	// 因为是公共前缀,循环次数为第一个字符串的长度即可
	for i := 0; i < len(strs[0]); i++ {
		// 取出第一个字符以外的字符
		for j := 1; j < len(strs); j++ {
			// strs[j][i] 第j个字符串的第i个字符 与 第一个字符串的第i个字符进行比较
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func main() {

}
