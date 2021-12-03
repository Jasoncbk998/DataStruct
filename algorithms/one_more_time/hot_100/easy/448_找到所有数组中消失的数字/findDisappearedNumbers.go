/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/23 9:39 下午
 **/
package main

/**
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，
数组中的元素一些出现了两次，另一些只出现一次。找到所有在 [1, n] 范围之间没有出现在数组中的数字
Input:
[4,3,2,7,8,2,3,1]
Output:
[5,6]
*/
/**
找出 [1,n] 范围内没有出现在数组中的数字。要求不使用额外空间，并且时间复杂度为 O(n)。
要求不能使用额外的空间，那么只能想办法在原有数组上进行修改，并且这个修改是可还原的。
时间复杂度也只能允许我们一层循环。只要循环一次能标记出已经出现过的数字，这道题就可以按要求解答出来。
这里笔者的标记方法是把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[| nums[i] |- 1] * -1。
这里需要注意的是，nums[i] 需要加绝对值，因为它可能被之前的数置为负数了，需要还原一下。
最后再遍历一次数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。把结果输出到最终数组里即可。
*/
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		//负数->正数
		if v < 0 {
			v = -v
		}
		// 存在则标记为负数
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
func findDisappearedNumbers2(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for _, v := range nums {
		v = (v - 1) % n
		//表示有,且不影响取余情况
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
