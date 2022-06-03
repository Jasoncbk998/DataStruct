/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/3 18:18
 **/
package main

import "fmt"

func main() {
	var whatever [3]struct{}

	for i := range whatever {
		defer func(i int) {
			fmt.Println(i) // 2,1,0
		}(i)
	}
	defer fmt.Println("---")
	for i := range whatever {
		defer func() {
			fmt.Println(i) // 2,2,2
		}()
	}
}
