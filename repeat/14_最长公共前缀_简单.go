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
	//因为是最长公共子串,所以外层循环,就是第一个元素的字节长度,纵向对比的思想一定要清晰
	for i := 0; i < len(strs[0]); i++ {
		//内层循环是数组长度
		for j := 1; j < len(strs); j++ {
			//	i == len(strs[j]) 说明i的长度至少等于j,所以可以return
			//  strs[j][i] != strs[0][i] 因为是纵向比对,第j个元素均和第一个元素比较i位
			//理解纵向比对,由于依次比对
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				// 因为是公共最长子串,所以数组的第一个元素中取长度就可以,最长的就是第一个元素本身
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix_(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// abc,abcdd,ab
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
