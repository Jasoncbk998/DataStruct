/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/11 2:49 下午
 **/
package main

/**
输入：s = "abcdefg", k = 2
输出："bacdfeg"
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
要求按照一定规则反转字符串：每 2 * K 长度的字符串，
反转前 K 个字符，后 K 个字符串保持不变；
对于末尾不够 2 * K 的字符串，如果长度大于 K，那么反转前 K 个字符串，剩下的保持不变。
如果长度小于 K，则把小于 K 的这部分字符串全部反转。
这一题是简单题，按照题意反转字符串即可。
*/
func reverseStr(s string, k int) string {
	if k > len(s) {
		k = len(s)
	}
	for i := 0; i < len(s); i = i + 2*k {
		if len(s)-i >= k {
			ss := revers(s[i : i+k])
			s = s[:i] + ss + s[i+k:]
		} else {
			ss := revers(s[i:])
			s = s[:i] + ss
		}
	}
	return s
}
func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
