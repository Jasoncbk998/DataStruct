/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/26 10:00 上午
 **/
package main

/**
给你两个 没有重复元素 的数组nums1 和nums2,其中nums1是nums2的子集。
请你找出 nums1中每个元素在nums2中的下一个比其大的值。
nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。如果不存在，对应位置输出 -1
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
    对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
    对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
*/

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	//先循环子集
	for i, v1 := range nums1 {
		start := false
		//子集的元素依次与父集的元素比较
		for _, v2 := range nums2 {
			if v2 == v1 {
				start = true
			}
			//寻找更大值,有则覆盖
			if start && v2 > v1 {
				nums1[i] = v2
				break
			}
		}
		//在nums2中没有更大值,则-1
		if nums1[i] == v1 {
			nums1[i] = -1
		}
	}
	return nums1
}
