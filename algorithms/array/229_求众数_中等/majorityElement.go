/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/22 9:42 下午
 **/
package main

/**
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
输入：[3,2,3]
输出：[3]
输入：[1,1,1,3,3,2,2,2]
输出：[1,2]
*/
// 利用hash表
func majorityElement(nums []int) []int {
	ans := []int{}
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for v, c := range cnt {
		if c > len(nums)/3 {
			ans = append(ans, v)
		}
	}
	return ans
}
