/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:07 下午
 **/
package main

import "fmt"

func main() {
	var ele int = 10
	//相当于右移取半
	//左移 翻倍
	//  5  像右移动一位 相当于把10 换算为2进制,然后向右移动一位 将一个数的各二进制位全部右移若干位，正数左补0，负数左补1，右边丢弃。
	//箭头像右,就是右移
	fmt.Println(ele >> 2)

	//像左移动一位
	//20
	//将一个运算对象的各二进制位全部左移若干位（左边的二进制位丢弃，右边补0）。
	fmt.Println(ele << 3)

}
