/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/18 9:29 下午
 **/
package main

/**
有序数组,给定一个元素,插入到该数组,返回对应位置
输入: nums = [1,3,5,6], target = 5
输出: 2
*/
/**
思想就是,二分查找,通过比对大小确定具体位置,也就是索引序号
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n //最后的位置
	for left <= right {
		mid := left + (right-left)>>1
		// 目标值在右边则移动左边界,反之移动右边界
		if target >= nums[mid] {
			ans = mid // 位置
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
