/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/21 9:53 下午
 **/
package main

import "fmt"

func main() {
	s := "egg"
	t := "add"
	isomorphic := isIsomorphic(s, t)
	fmt.Println(isomorphic)
}

/**
s = "egg", t = "add"

输入：s = "egg", t = "add"
输出：true
示例 2：
输入：s = "foo", t = "bar"
输出：false

做一种映射,比如 e只对应a  如果e对应到其他字符 则认为不是同构

*/
func isIsomorphic(s string, t string) bool {
	s1 := map[byte]byte{}
	s2 := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s1[x] > 0 && s1[x] != y || s2[y] > 0 && s2[y] != x {
			return false
		}
		s1[x] = y
		s2[y] = x
	}
	return true
}
func isIsomorphic1(s string, t string) bool {
	s1 := []rune(s)
	s2 := []rune(t)
	pairA := make(map[rune]rune)
	pairB := make(map[rune]rune)
	for i := 0; i < len(s1); i++ {
		a, b := s1[i], s2[i]
		if (pairA[a] > 0 && pairA[a] != b) || (pairB[b] > 0 && pairB[b] != a) {
			return false
		}
		pairA[a] = b
		pairB[b] = a
	}
	return true
}
