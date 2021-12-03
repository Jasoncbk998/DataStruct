/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/18 2:15 下午
 **/
package merge_sort

import "fmt"

/**
归并排序
讲一个序列,不断地分割为两个子序列.然后将不能分割的多个子序列,进行合并
分割-->合并



*/
func Sort(arr []int) {
	//创建int数组
	var s = make([]int, len(arr)/2+1)
	if len(arr) < 2 {
		return
	}
	//找到中间位置
	mid := len(arr) / 2
	//左边切割
	Sort(arr[:mid])
	fmt.Println(arr)
	//右边切割
	Sort(arr[mid:])
	//左 小于右
	if arr[mid-1] <= arr[mid] {
		return
	}

	copy(s, arr[:mid])
	l, r := 0, mid
	for i := 0; ; i++ {
		if s[l] <= arr[r] {
			arr[i] = s[l]
			l++
			if l == mid {
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) {
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}
	return
}
