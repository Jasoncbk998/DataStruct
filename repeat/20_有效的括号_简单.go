/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/22 12:45 下午
 **/
package main

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
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
	if n%2 == 1 {
		return false
	}
	//map是 )(,}{,][
	paris := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		//有
		if paris[s[i]] > 0 {
			if paris[s[i]] == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			//无
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
