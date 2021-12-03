/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/17 1:21 下午
 **/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ints := []int{-1, 2, 2, 1, -4}
	fmt.Println(threeSumClosest(ints, 1))
}

/**
Given array nums = [-1, 2, 1, -4], and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	// 必须3个数
	if n > 2 {
		sort.Ints(nums)
		// i从0取到n-2 n-1留给k
		// j是i的后一位
		// 通过这样 安排 i j k 去遍历全部值
		// i每次确定后,开始确定j k 然后固定i 让j k进行遍历所有组合,进行匹配
		for i := 0; i < n-2; i++ {
			//相邻两位且等值,略过 这个是为了提升效率,不用多次计算同一个值,属于优化手段
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				//小于diff
				if abs(sum-target) < diff {
					// 更新diff
					res, diff = sum, abs(sum-target)
				}
				//进行夹逼
				if sum == target {
					return res
				} else if sum > target {
					k--
					//j++
				} else {
					j++
					//k--
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
