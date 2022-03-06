/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/21 12:48 下午
 **/
package main

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		// 遍历按照比对规则去map中 获取
		// map中无此元素则put,有则弹出,通过比对stack长度判断是否是有效的括号
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
