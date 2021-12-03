/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 12:23 下午
 **/
package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum2_test(nums, target))

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if _, ok := m[tmp]; ok {
			return []int{m[tmp], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// nums:=[]int{2,7,11,15}
//target:=18
func twoSum1(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		//map中存储的是每一个元素,如果发现,tmp等于数组中的元素,那就说明找到
		//反之遍历像map中插入,k是nums[i],v是索引
		if _, ok := maps[tmp]; !ok {
			return []int{i, maps[tmp]}
		}
		maps[nums[i]] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if _, ok := maps[tmp]; !ok {
			return []int{i, maps[tmp]}
		}
		maps[nums[i]] = i
	}
	return nil
}

func twoSum2_test(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))
	for k, _ := range nums {
		temp := target - nums[k]
		if _, ok := maps[temp]; ok {
			return []int{k, maps[temp]}
		}
		//map的key是值,value是索引
		maps[nums[k]] = k
	}
	return nil
}

/**
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
*/
func twoSum2_test111(nums []int, target int) []int {

	//map[value]index
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
