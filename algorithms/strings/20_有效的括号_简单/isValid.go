/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/15 5:24 下午
 **/
package main

import "fmt"

func main() {
	s := "()[]{}"
	valid := isValid(s)
	fmt.Println(valid)
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
输入：s = "()"
输出：true
输入：s = "()[]{}"
输出：true
*/
func isValid(s string) bool {
	n := len(s)
	// 因为括号是成对出现的,所以是奇数就肯定错
	if n%2 == 1 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	//从string的0位开始
	//s = "()[]{}" true
	// s = "{[]}" true
	for i := 0; i < n; i++ {
		//模拟栈,因为把括号存储在 map中,而且是成对出现的括号 ,那么不用考虑先后,只要匹配成功就--
		if paris[s[i]] > 0 {
			//判断栈
			if len(stack) == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			//如果匹配,则弹出,栈是一个数组,发现匹配就remove该元素
			stack = stack[:len(stack)-1]
		} else {
			//追加
			stack = append(stack, s[i])
		}
	}
	//最终判断栈的长度是否等于0
	return len(stack) == 0
}
