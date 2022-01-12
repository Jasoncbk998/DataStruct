/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/16 6:33 下午
 **/
package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{5, 4, 3, 2, 1}
	sort := selectSort(nums)
	fmt.Println(sort)
	arr := []string{"2c", "1a", "1b", "1x", "1z", "1m", "1n", "1d", "1f"}

	sort_string := selectSort_string(arr)
	fmt.Println(sort_string)

}
func selectSort_string(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := 1; j < length; j++ {
				//字符串排序
				if strings.Compare(arr[min], arr[j]) > 0 {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
		return arr
	}

}
func selectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
		fmt.Println(arr)
	}
	return arr
}

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}

}
