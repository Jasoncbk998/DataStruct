/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 11:15 上午
 **/
package main

import "fmt"

func main() {
	paris := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	b := paris['a']
	fmt.Printf("%t\t%v", b, b)
}
