/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/7 12:46 下午
 **/
package main

import (
	"math/rand"
	"time"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
感觉还是直接排序,然后返回 最好
*/
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left, right, index int) int {
	p := partition(nums, left, right)
	if p == index {
		return nums[p]
	} else if p < index {
		return quickSort(nums, p+1, right, index)
	}
	return quickSort(nums, left, p-1, index)
}

func partition(nums []int, left, right int) int {
	p := rand.Int()%(right-left+1) + left
	nums[right], nums[p] = nums[p], nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= nums[right] {
			nums[i], nums[j] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
