/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 12:23 下午
 **/
package main

/**
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
*/
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := maps[temp]; ok {
			return []int{maps[temp], i}
		}
		maps[nums[i]] = i
	}
	return nil
}

func test(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := maps[temp]; ok {
			return []int{maps[temp], i}
		}
		maps[nums[i]] = i
	}
	return nil
}
