/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/10 2:17 下午
 **/
package main

import "fmt"

/**
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
把数组一的每个数字都存进字典中，然后在数组二中依次判断字典中是否存在，如果存在，在字典中删除它(因为输出要求只输出一次)
*/
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	//nums1 加入到字典中
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		fmt.Println(n, m[n])
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
