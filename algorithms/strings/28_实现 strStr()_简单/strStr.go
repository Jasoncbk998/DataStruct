/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/18 5:30 下午
 **/
package main

import "fmt"

func main() {
	str := strStr("hello", "ll")
	fmt.Println(str)

}

//  实现一个查找 substring 的函数。
// 如果在母串中找到了子串，返回子串在母串中出现的下标，如果没有找到，返回 -1，如果子串是空串，则返回 0 。
// haystack母串   needle子串
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			//开始就找到了
			if j == len(needle) {
				return i
			}
			//越界
			if i+j == len(haystack) {
				return -1
			}
			//子串中从第一位开始找,发现相等继续在子循环中寻找,发现不相等直接跳出内层循环
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
