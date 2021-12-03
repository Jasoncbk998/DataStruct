/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/26 10:00 上午
 **/
package main

import "fmt"

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

func main() {

	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	element := nextGreaterElement1(nums1, nums2)
	fmt.Println(element)

}

//单调栈 + 哈希表
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res

}
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	//先循环子集
	for i, v1 := range nums1 {
		start := false
		//子集的元素依次与父集的元素比较
		for _, v2 := range nums2 {
			if v2 == v1 {
				start = true
			}
			if start && v2 > v1 {
				nums1[i] = v2
				break
			}
		}
		if nums1[i] == v1 {
			nums1[i] = -1
		}
	}
	return nums1
}
