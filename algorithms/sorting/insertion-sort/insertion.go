/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/17 7:04 下午
 **/
package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sort(arr)
	fmt.Println(arr)
}

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
		cur := arr[i]
		prev := i - 1
		//局部
		for prev >= 0 && arr[prev] > cur {
			arr[prev+1] = arr[prev]
			prev = prev - 1
		}
		arr[prev+1] = cur
	}
}

func test(arr []int) {
	// 5,4,3,2,1
	for i := 1; i < len(arr); i++ {
		//4
		cur := arr[i]
		//0
		prev := i - 1
		// 5>4
		for prev >= 0 && arr[prev] > cur {
			//arr[1]=arr[0]
			//
			arr[prev+1] = arr[prev]
			prev = prev - 1
		}
		arr[prev+1] = cur
	}

}
