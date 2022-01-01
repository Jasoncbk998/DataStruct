/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/17 3:59 下午
 **/
package main

import "fmt"

/**
查找目标值的索引
分段的二分查找,因为是一个有序数组,按照某一个元素进行反转,也就是说在数组中两部分有序
也就是局部有序,还是二分查找的思路
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 认为左边是low 是较小的,右边是high,是较大的
	low, high := 0, len(nums)-1
	for low <= high {
		//不断逼近
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			//中心思想
			//因为是一个顺序数组,按照某一个元素进行反转,所以不管以为为轴,进行反转 最左边和最右边一定在各自的局部有序数组内部
			//nums = [4,5,6,7,0,1,2], target = 0
			//当目标值大于较小值,并且目标值小于中间值,进行夹逼
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 5
	i := search(nums, target)
	fmt.Println(i)

}
