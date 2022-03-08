/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/8 11:36 上午
 **/
package main

/**
输入：strs = ["flower","flow","flight"]
输出："fl"
纵向扫描
*/
//str:=[]string{"flower","flow","flight"}
//u := str[0][0]
//fmt.Println(string(u))// f
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//因为是最长公共子串,所以外层循环,就是第一个元素的字节长度
	for i := 0; i < len(strs[0]); i++ {
		//内层循环是数组长度
		for j := 1; j < len(strs); j++ {
			//i == len(strs[j]) 说明i的长度至少等于j,所以可以return
			//  strs[i][j] != strs[0][i] 是纵向比对,数组中元素的j位字符
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
