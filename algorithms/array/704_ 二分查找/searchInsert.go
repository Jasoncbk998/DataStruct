/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 1:26 下午
 **/
package _04__二分查找

func SearchInsert(nums []int, target int) int {
	//数组长度
	n := len(nums)
	//划定边界
	left, right := 0, n-1
	// 记录索引位置
	ans := n
	// 夹逼定理的思想
	for left <= right {
		// 数组的中间索引
		mid := (right-left)>>1 + left
		//目标值<=数组中建索引
		if target <= nums[mid] {
			//  更新索引位置
			ans = mid
			// 从mid开始左移一位
			right = mid - 1
		} else {
			// 如果target>nums[mid]
			// 说明,target位于(mid+1,right]中,所以left右移mid+1
			left = mid + 1
		}
	}
	return ans
}
