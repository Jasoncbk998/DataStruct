/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/1 8:49 下午
 **/
package main

func isAnagram(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}
