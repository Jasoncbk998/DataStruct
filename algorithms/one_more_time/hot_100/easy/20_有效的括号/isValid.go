/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 12:30 下午
 **/
package main

/**
Input: "([)]"
Output: false

Input: "{[]}"
Output: true
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if (v == ']') && len(stack) > 0 && stack[len(stack)]-1 == '[' ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
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
		//模拟栈
		//在map中可以根据key取到对应value,则返回byte大于0
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

func isValid2_test1(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 1; i < n; i++ {
		if paris[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func test(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 1; i < n; i++ {
		if paris[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != paris[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
