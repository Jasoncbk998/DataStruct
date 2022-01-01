/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/18 7:02 下午
 **/
package main

func main() {

}

/**
有序数组,给定一个元素,插入到该数组,返回对应位置
输入: nums = [1,3,5,6], target = 5
输出: 2
*/
// 二分查找,很简单
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
