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
题目给出 2 个数组 A 和 B，针对 A 中的每个数组中的元素，
要求在 B 数组中找出比 A 数组中元素大的数
B 中元素之间的顺序保持不变。如果找到了就输出这个值，如果找不到就输出 -1。
*/

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res, record := []int{}, map[int]int{}
	for i, v := range nums2 {
		record[v] = i
	}
	for i := 0; i < len(nums1); i++ {
		flag := false
		for j := record[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				flag = true
				break
			}
		}
		if flag == false {
			res = append(res, -1)
		}
	}
	return res
}
