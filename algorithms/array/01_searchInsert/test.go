/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/15 11:04 下午
 **/
package _1_searchInsert

func SearchInsert1(nums []int, target int) int {
	length := len(nums)
	index := length
	left, right := 0, length-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] <= target {
			left = mid + 1
		} else {
			index = mid
			right = mid - 1
		}
	}
	return index
}
