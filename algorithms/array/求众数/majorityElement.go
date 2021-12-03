/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/22 9:42 下午
 **/
package main

import "fmt"

/**
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
输入：[3,2,3]
输出：[3]

输入：[1,1,1,3,3,2,2,2]
输出：[1,2]

输入：nums = [1]
输出：[1]

*/
func main() {

	ints := []int{1, 1, 1, 3, 2, 2, 2}
	fmt.Println(len(ints), len(ints)/3)
	element := majorityElement(ints)
	fmt.Println(element)
}

// 利用hash表
func majorityElement(nums []int) []int {
	ans := []int{}
	cnt := map[int]int{}
	for _, v := range nums {
		//元素是value,key是次数
		cnt[v]++
	}
	for v, c := range cnt {
		if c > len(nums)/3 {
			ans = append(ans, v)
		}
	}
	return ans
}

type poll struct {
	name int
	vote int
}

func majorityElement1(nums []int) []int {
	a, b := poll{name: nums[0]}, poll{name: nums[0]}
	for _, num := range nums {
		// equal
		if a.name == num || b.name == num {
			if a.name == num {
				a.vote++
			}
			if b.name == num && b.name != a.name {
				b.vote++
			}
		} else { //not equal
			if a.name == b.name { // initiate b
				b.name = num
				b.vote = 1
			} else {
				a.vote--
				b.vote--
				if b.vote < 0 {
					b.name, b.vote = num, 1
					a.vote++
				}
			}
		}
		if a.vote < b.vote {
			a, b = b, a
		}
	}
	res := []int{}
	cnta, cntb := 0, 0
	for _, num := range nums {
		if num == a.name {
			cnta++
		}
		if num == b.name {
			cntb++
		}
	}
	if cnta > len(nums)/3 {
		res = append(res, a.name)
	}
	if cntb > len(nums)/3 && b.name != a.name {
		res = append(res, b.name)
	}
	return res
}
