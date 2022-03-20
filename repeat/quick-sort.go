/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/7 1:13 下午
 **/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{2, 3, 5, 58, 1, 3}
	fmt.Println(slice)
	ints := sort2(slice)
	fmt.Println(ints)
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

func sort3(nums []int) []int {
	// 荷兰国旗问题
	netherlandsFlag := func(nums []int, l, r int) (start, end int) {
		if l >= r {
			return -1, -1
		}
		// 区间[l, less] 为小于主元的元素
		less := l - 1
		//区间[more,r]为大于主元的元素
		more := r
		i := l
		//与大于区相遇时停止遍历
		for i < more {
			if nums[i] < nums[r] {
				//小于区右扩,交换元素,i右移
				less++
				nums[less], nums[i] = nums[i], nums[less]
				i++
			} else if nums[i] > nums[r] {
				//大于区左扩,交换元素,i不动
				more--
				nums[i], nums[more] = nums[more], nums[i]
			} else if nums[i] == nums[r] {
				//i 右移
				i++
			}
		}
		//将主元与大于区第一个元素交换
		nums[more], nums[r] = nums[r], nums[more]
		return less + 1, more
	}
	//插入排序
	insertSort := func(nums []int) {
		length := len(nums)
		for i := 1; i < length; i++ {
			temp := nums[i]
			j := i
			for j > 0 && nums[j-1] > temp {
				nums[j] = nums[j-1]
				j--
			}
			nums[j] = temp
		}
	}
	var quickSort func(nums []int, l, r int)
	quickSort = func(nums []int, l, r int) {
		if l >= r {
			return
		}
		//规模小于100 用插入排序
		if r-l < 100 {
			insertSort(nums[l : r+1])
			return
		}
		p := rand.Intn(r-l+1) + l
		nums[p], nums[r] = nums[r], nums[p]
		start, end := netherlandsFlag(nums, l, r)
		quickSort(nums, l, start-1)
		quickSort(nums, end+1, r)
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}
