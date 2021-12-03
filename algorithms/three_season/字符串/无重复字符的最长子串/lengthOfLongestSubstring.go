/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/13 8:37 下午
 **/
package main

import "fmt"

func main() {
	s := "aaaaaa"
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return -1
	}
	chars := []byte(s)
	if len(chars) == 0 {
		return -1
	}
	prevIdxes := make([]int, 128)
	for i := 0; i < len(prevIdxes); i++ {
		prevIdxes[i] = -1
	}
	prevIdxes[chars[0]] = 0
	// 以i - 1位置字符结尾的最长不重复字符串的开始索引（最左索引）
	li := 0
	max := 1
	for i := 1; i < len(chars); i++ {
		pi := prevIdxes[chars[i]]
		if li <= pi {
			li = pi + 1
		}
		//存储这个字符出现的位置
		prevIdxes[chars[i]] = i
		max = Max(max, i-li+1)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
