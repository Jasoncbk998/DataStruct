/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/13 4:27 下午
 **/
package main

import "fmt"

/**
输入：s = "the sky is blue"
输出："blue is sky the"

输入：s = " hello world "
输出："world hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包括。

输入：s = "a good  example"
输出："example good a"
解释：如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。

输入：s = "  Bob    Loves  Alice   "
输出："Alice Loves Bob"

输入：s = "Alice does not even like bob"
输出："bob like even not does Alice"
*/
// https://leetcode-cn.com/problems/reverse-words-in-a-string/
func main() {
	s := "   is good example"
	words := reverseWords(s)
	fmt.Println(words)

}

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	cur := 0
	//字符串最终的有效长度
	length := 0
	space := true
	chars := []byte(s)
	//去除首尾空格,并移动byte
	for i := 0; i < len(chars); i++ {
		//非空格
		if chars[i] != ' ' {
			chars[cur] = chars[i]
			cur++
			space = false
		} else if space == false {
			// 空格
			chars[cur] = ' '
			cur++
			space = true
		}
	}
	//确定有效长度
	if space {
		length = cur - 1
	} else {
		length = cur
	}
	if length <= 0 {
		return ""
	}
	//整体逆序
	reverse(chars, 0, length)
	// 对每一个单词进行逆序
	//前一个空格字符的位置
	prevSpaceIdx := -1
	// i<length 因为是移动原string,后边可能有残留的字符,所以不能等于
	for i := 0; i < length; i++ {
		//确定空格索引位置
		if chars[i] != ' ' {
			continue
		}
		//发现空格进行反转
		reverse(chars, prevSpaceIdx+1, i)
		prevSpaceIdx = i
		fmt.Println(prevSpaceIdx)
	}
	//反转最后一个单词
	//[ ) 左闭右开
	reverse(chars, prevSpaceIdx+1, length)
	return string(chars[:length])
}

//[right,left] 局部逆序
func reverse(chars []byte, right int, left int) {
	left--
	for right < left {
		chars[right], chars[left] = chars[left], chars[right]
		right++
		left--
	}
}
