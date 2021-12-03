/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/27 9:26 下午
 **/
package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}
	intersect := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersect = append(intersect, num)
			m[num]--
		}
	}
	return intersect

}
