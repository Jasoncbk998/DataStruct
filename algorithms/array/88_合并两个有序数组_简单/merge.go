/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/14 10:22 下午
 **/
package main

/**
合并两个已经有序的数组，结果放在第一个数组中，第一个数组假设空间足够大。要求算法时间复杂度足够低。
为了不大量移动元素，
因为是两个递增的顺序数组,所以从先比较最大的,放到nums1最后边,依次摆放
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	//对两个数组进行排序
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	//这种情况是
	//nums1: 4,5,6,0,0,0 m=3
	//nums2: 1,2,3 n=3
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
