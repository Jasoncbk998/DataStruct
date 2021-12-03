/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/14 10:22 下午
 **/
package main

import "fmt"

/**
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/
func main() {
	ints := []int{2, 4, 6, 7, 8, 9}
	i := []int{1, 3, 5}
	merge1(ints, len(ints), i, len(i))
	fmt.Println(ints)
}

//两个数组都是非递减顺序数组
//双指针
// m是nums1的长度 n是nums2的长度
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	// 双指针
	p1, p2 := 0, 0
	//开始循环
	for {
		//每次进入循环都要判断,如果一个数组已经比较完成那么指针就会等于长度,则跳出for
		if p1 == m {
			//将nums2赋值到sorted中,等待合并
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p2:]...)
			break
		}
		//两个数组的元素进行比对,根据两个指针,不断地比较,进行追加到sorted
		if nums1[p1] < nums2[p2] {
			// nums1追加到sorted
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	//将sorted排序后的有序数组拷贝到nums1
	copy(nums1, sorted)
}
