/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/27 12:31 下午
 **/
package main

import (
	"fmt"
	"sort"
)

/**
三个数相加为0
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
/**
用 map 提前计算好任意 2 个数字之和，保存起来，可以将时间复杂度降到 O(n^2)。
这一题比较麻烦的一点在于，最后输出解的时候，要求输出不重复的解。
数组中同一个数字可能出现多次，同一个数字也可能使用多次，但是最后输出解的时候，不能重复。
例如 [-1，-1，2] 和 [2, -1, -1]、[-1, 2, -1] 这 3 个解是重复的，
即使 -1 可能出现 100 次，每次使用的 -1 的数组下标都是不同的。

这里就需要去重和排序了。
map 记录每个数字出现的次数，然后对 map 的 key 数组进行排序，
最后在这个排序以后的数组里面扫，找到另外 2 个数字能和自己组成 0 的组合。
*/
func main() {
	nums := []int{1, -1, 0}
	sum2 := threeSum2(nums)
	fmt.Println(sum2)
}

/**
穷尽枚举
*/
func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 枚举a,不断循环
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		//确定c的索引,通过挪动a和b的索引,来完成三个数字的匹配
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				//进入下一次循环
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
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
