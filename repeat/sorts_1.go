/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/4/17 18:20
 **/
package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	InsertionSort(arr)
	fmt.Println(arr)
}

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	/**
	a [5,4,3,2,1] n=5
	i=0 j<4
	i=1 j<3
	i=2 j<2
	i=3 j<1
	i=4 j<0
	*/
	for i := 0; i < n; i++ {
		//提前退出标志
		falg := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				// 此次冒泡排序有数据交换
				falg = true
			}
		}
		// 如果没有交换数据 提前退出
		if !falg {
			break
		}
	}
}

/**
插入排序
核心思想就是,保证左边的小区间是有序的
原数组:4,5,6,1,3,2
第一次插入排序: 4 ,5,6,1,3,2
第二次插入排序: 4,5, 6,1,3,2
第三次插入排序: 4,5,6 ,1,3,2
第四次插入排序: 1,4,5,6 ,3,2
1,3,4,5,6 ,2
移动元素进行比较,然后插入
*/
func InsertionSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		//取元素
		value := a[i]
		// 左边,小区间比较范围
		j := i - 1
		//开始小区间内保证局部有序
		for ; j >= 0; j-- {
			// 开始比较,每次value都与小区间内的最大值进行比较
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		// 平移后把最左侧元素赋值
		a[j+1] = value
	}
}

//选择排序
// 核心思路和插入排序很像,小范围找到最小值的索引值,然后交换,不断缩小比较范围,最后从局部有序到整体有序
func SelectionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		//查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
