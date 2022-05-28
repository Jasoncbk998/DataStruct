/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/26 11:53
 **/
package main

import "sort"

/**
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次
注意：解集不能包含重复的组合。
输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	// c是子数组
	// res 是最终结果,二维数组
	c, res := []int{}, [][]int{}
	sort.Ints(candidates) // 排序
	findcombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	//每次递归进到这里,进行target比较,等于0意味着寻找完毕
	if target == 0 {
		news := make([]int, len(c))
		copy(news, c)
		*res = append(*res, news)
		return
	}
	for i := index; i < len(nums); i++ {
		// 去重
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findcombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
