/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/16 9:22 下午
 **/
package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1}
	sort_test(nums)
	fmt.Println(nums)
}

/**
每一次都会找到一个最大值,放到最后
通过两个for
第一个for 是--
第二个for 是++
第二个for是实际进行比较,并且进行换位,比较完以后,在当前的数字中找到一个最大值,放到最后,跳出循环,然后第一个for -- 在继续 j<i
*/
func sort(arr []int) {
	//遍历一遍
	for itemCount := len(arr) - 1; ; itemCount-- {
		swap := false
		for i := 1; i <= itemCount; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

func sort_test(arr []int) {
	for i := len(arr) - 1; ; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
