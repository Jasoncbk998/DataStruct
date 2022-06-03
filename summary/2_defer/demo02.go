/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/6/3 18:28
 **/
package main

import "fmt"

type number int

func (n number) print() {
	fmt.Println(n)
}
func (n *number) pprint() {
	fmt.Println(*n)
}
func main1() {
	var n number
	defer n.print()
	defer n.pprint()
	defer func() {
		n.print()
	}()
	// 闭包
	defer func() {
		n.pprint()
	}()
	n = 3
	/**
	  3
	  3
	  3
	  0
	*/

}

func main() {
	r := f()
	fmt.Println("main:", r) // 5
	fmt.Println("---")
	r2 := f2()
	fmt.Println("main:", r2) // 5
	r3 := f3()
	fmt.Println("main:", r3) // 0
}
func f2() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return r
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return r
}

func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println("defer -t :", t)   // 10
		fmt.Println("defer  -r  :", r) // 5
	}()
	fmt.Println("f方法-t:", t) // 5
	fmt.Println("f方法-r:", r) //  0
	return t
}
