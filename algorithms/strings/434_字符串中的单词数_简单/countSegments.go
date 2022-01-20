/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/19 1:15 下午
 **/
package main

import "fmt"

func main() {
	segments := countSegments("aaa   dd d d ")
	fmt.Println(segments)
}

/**
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
请注意，你可以假定字符串里不包括任何不可打印的字符
*/
//按照每个字符进行遍历,发现空格就++
func countSegments(s string) int {
	//两个单词中多个空格认为是一个空格
	segments := false
	cnt := 0
	for _, v := range s {
		if v == ' ' && segments {
			//遇到空格
			segments = false
			cnt += 1
		} else if v != ' ' {
			//非空格
			segments = true
		}
	}
	if segments {
		cnt++
	}
	return cnt
}
