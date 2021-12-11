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
			//开始寻找
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
