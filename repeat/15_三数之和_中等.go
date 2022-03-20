/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/19 4:00 下午
 **/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, -3, 3, -1, -4, 4, 5}
	sum := threeSum(nums)
	fmt.Println(sum)
}

/**
三数之和等于0
第二重循环枚举到的元素不小于当前第一重循环枚举到的元素；

第三重循环枚举到的元素不小于当前第二重循环枚举到的元素。
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 枚举a
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 保证和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要报咋整b的指针在c指针左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//如果指针重合,随着b后续的增加
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
