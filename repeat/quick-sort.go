/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/7 1:13 下午
 **/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{2, 3, 5, 58, 1, 3}
	fmt.Println(slice)
	ints := sort2(slice)
	fmt.Println(ints)
}
func sort(arr []int, low, high int) []int {
	if high > low {
		privot := part(arr, low, high)
		sort(arr, low, privot-1)
		sort(arr, privot+1, high)
	}
	return arr
}

func part(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && pivot <= arr[high] {
			high--
		}
		arr[low] = arr[high]
		for low < high && pivot >= arr[low] {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func getRateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func sort2(arr []int) []int {
	var recurse func(left, right int)
	var partition func(left, right, pivot int) int

	partition = func(left, right, pivot int) int {
		v := arr[pivot]
		right--
		arr[pivot], arr[right] = arr[right], arr[pivot]
		for i := left; i < right; i++ {
			if arr[i] <= v {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}
		arr[left], arr[right] = arr[right], arr[left]
		return left
	}

	recurse = func(left, right int) {
		if left < right {
			pivot := (right + left) / 2
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}
	recurse(0, len(arr))
	return arr
}
