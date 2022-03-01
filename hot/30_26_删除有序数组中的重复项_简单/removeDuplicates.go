/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/1 4:29 下午
 **/
package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 1 1 2 2 3
	// 1 2 3 4
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
