/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/17 7:04 下午
 **/
package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sort2(arr)
	fmt.Println(arr)
}

/**
多数文件排序,插入排序
*/
/**
基本思路就是
拿到一个数组,用他的后位比较前位,如果发现后位小于前位,则进行换位
比如 5,4,3,2,1
i=1 5>4 则 4,4,3,2,1 value=5 arr[j+1]=arr[1]=value=5 -> 4,5,3,2,1
*/
func sort(arr []int) {
	//每次比较(0,i)中的最大值,放到i的位置
	//局部有序->整体有序
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		//后位和前位比,后位小于前位则交换,j-1
		//若 f(n)<f(n-1) 则 换位,n=n-1 将value
		// f(n-1)<f(n-2) 则 换位 n=n-2
		//n-1位
		//保证可以循环到最前边
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = value
	}
}
func sort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		prev := i - 1
		for prev >= 0 && arr[prev] > cur {
			arr[prev+1] = arr[prev]
			prev = prev - 1
		}
		arr[prev+1] = cur
	}
}
