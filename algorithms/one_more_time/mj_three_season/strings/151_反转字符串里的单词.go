/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/29 12:30 下午
 **/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the     sky    is blue"
	println(reverseWords(s))
}

func reverseWords151(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}
func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	res := []byte(s)
	length := len(res)
	slow, fast := 0, 0

	//去除头部冗余空格
	for fast < length && res[fast] == ' ' {
		fast++
	}

	//去除单词间冗余空格
	for ; fast < length; fast++ {
		if fast-1 > 0 && res[fast-1] == res[fast] && res[fast] == ' ' {
			continue
		} else {
			res[slow] = res[fast]
			slow++
		}
	}

	//去除尾部冗余空格
	if slow-1 > 0 && res[slow-1] == ' ' {
		res = res[:slow-1]
	} else {
		res = res[:slow]
	}

	//总体反转字符串
	reverse(res, 0, len(res)-1) //eulb si yks eht

	//反转单个单词
	i := 0
	for i < len(res) {
		j := i
		//寻找第一个单词的截止点,通过j记录,每次更新j
		for ; j < len(res) && res[j] != ' '; j++ {
		}
		//左闭右开
		reverse(res, i, j-1)
		fmt.Println("i:", string(res))
		i = j // 下一个单词的起始点
		i++
	}
	return string(res)
}

func reverse(res []byte, left, right int) {
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
}
