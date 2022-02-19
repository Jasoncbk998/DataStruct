/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/10 12:57 下午
 **/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//slice := generateSlice(5)
	slice := []int{5, 6, 3, 2, 1}
	fmt.Println(slice)

	ints := sort(slice, 0, len(slice)-1)
	fmt.Println(ints)
}

/**
从数列中挑出一个元素,称为 “基准”(pivot)
重新排序数列,所有元素比基准值小的摆放pivot左边,所有元素比基准值大的摆在pviot右边.
在这个分区退出之后,该基准就处于数列的中间位置.这个称为分区操作;
递归地(recursive)把小于基准值元素的子数列和大于基准值元素的子数列排序;
*/
func sort(arr []int, low, high int) []int {
	if high > low {
		//每一次循环,都在子区间内找到最大最小值,
		privot := partition(arr, low, high)
		sort(arr, low, privot-1)
		sort(arr, privot+1, high)
	}
	return arr
}

//极大值 靠右放 极小值 靠左放
func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]
		for low < high && pivot >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
