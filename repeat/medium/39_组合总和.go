/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/5/30 21:51
 **/
package main

/**
给你一个 无重复元素 的整数数组candidates 和一个目标整数target,找出candidates中可以使数字和为目标数target 的 所有不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为target 的不同组合数少于 150 个。
*/
/**
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
*/
var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	find(candidates, target, []int{}, 0)
	return res
}

func find(candidates []int, target int, oneRes []int, start int) {
	sum := 0
	for _, val := range oneRes {
		sum += val
	}
	if sum > target {
		return
	}
	if sum == target {
		copyArr := make([]int, len(oneRes))
		copy(copyArr, oneRes)
		res = append(res, copyArr)
		return
	}

	for i := start; i < len(candidates); i++ {
		oneRes = append(oneRes, candidates[i])
		find(candidates, target, oneRes, i) //再从i位置开始找可能组合
		oneRes = oneRes[:len(oneRes)-1]
	}
}
