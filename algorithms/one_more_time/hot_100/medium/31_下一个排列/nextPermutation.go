/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/24 12:04 下午
 **/
package main

import "fmt"

func main() {

	nums := []int{2, 1, 1}
	nextPermutation1(nums)
	fmt.Println(nums)

}

/**
Input: nums = [3,2,1]
Output: [1,2,3]
现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
*/

func nextPermutation1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// find: A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 { // 不是最后一个排列
		// find: A[i]<A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap A[i], A[k]
		nums[i], nums[k] = nums[k], nums[i]
	}
	// reverse A[j:end]
	// j开始到最后一个元素
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 比如 4,3,1,2 => 4,3,2,1 就可以 变成更大的一个数组
func nextPermutation_test(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// i 倒数第1个开始
	// j k 都是最后
	i, j, k := len(nums)-1, len(nums)-2, len(nums)-1
	//1,2,3,4
	// i 3	2	1	i=0
	// j 4 	3	2	j=1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// i=0
	// j=1
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j+1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
func nextPermutation(nums []int) {
	i, j := 0, 0
	for i := len(nums) - 2; i >= 0; i-- {
		// 前位小于后位,更新i
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			//更新j
			if nums[j] < nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i int, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
