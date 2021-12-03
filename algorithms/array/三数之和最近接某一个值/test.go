/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/17 6:48 下午
 **/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ints := []int{-1, 2, 2, 1, -4}
	fmt.Println(threeSumClosest4(ints, 1))
}

func threeSumClosest4(nums []int, target int) int {
	n := len(nums)
	if n < 2 {
		return -1
	}
	sort.Ints(nums)
	res := 0
	diff := math.MaxInt32
	for i := 0; i < n-2; i++ {
		//if i > 0 && nums[i] == nums[i-1] {
		//	continue
		//}
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if abs1(sum-target) < diff {
				res, diff = sum, abs1(sum-target)
			}
			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return res

}
func abs1(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
