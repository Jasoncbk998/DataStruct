/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 12:20 下午
 **/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	str := largestNumber(nums)
	fmt.Println(str)
}

/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
输入：nums = [10,2]
输出："210"
输入：nums = [3,30,34,5,9]
输出："9534330"
*/
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y < sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
