/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/19 4:30 下午
 **/
package main

import "fmt"

/**
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/
func main() {
	i := permute([]int{1, 2, 3})
	fmt.Println(i)
}
func permute(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	var p func(k int)
	//前k个元素的全排列
	p = func(k int) {
		if k == 1 {
			item := make([]int, n)
			copy(item, nums)
			result = append(result, item)
			return
		}
		//依次把1~k的元素放到最后并求k-1个元素全排列
		//然后归位
		for i := 0; i < k; i++ {
			nums[i], nums[k-1] = nums[k-1], nums[i]
			p(k - 1)
			nums[i], nums[k-1] = nums[k-1], nums[i]
		}
	}
	p(n)
	return result
}
