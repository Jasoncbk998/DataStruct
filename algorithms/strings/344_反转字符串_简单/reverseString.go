/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/28 8:45 下午
 **/
package main

/**
输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
*/

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		if left == right {
			break
		}
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
