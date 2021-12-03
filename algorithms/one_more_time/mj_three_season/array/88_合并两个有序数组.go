/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/3 10:39 上午
 **/
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 5, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 4, nums2, 3)
	fmt.Println(nums1)
}

/**
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			//放入p中
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
