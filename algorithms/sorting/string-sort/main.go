/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 6:52 下午
 **/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("ba", "bc"))
	fmt.Println(strings.Compare("c", "b"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a2", "a1"))
	fmt.Println(strings.Compare("a2", "a2"))

	fmt.Println("a" > "b")
	fmt.Println("a" < "b")
	fmt.Println("a" == "a")

	pb := "a1"
	pa := "a2"
	fmt.Println("pa", &pa) //go1.1 ，1.3比较地址，go1.10 优化，比较字符串
	fmt.Println("pb", &pb)
	fmt.Println(pa < pb)
	fmt.Println("a1" < "a2")
	fmt.Println("ab" < "ac")
	fmt.Println("a1b" < "a1c")
}
