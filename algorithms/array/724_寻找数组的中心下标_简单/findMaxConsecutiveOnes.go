/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/1/11 2:25 下午
 **/
package main

/**
在数组中，找到一个数，使得它左边的数之和等于它的右边的数之和，如果存在，则返回这个数的下标索引，否作返回 -1。
这里面存在一个等式，只需要满足这个等式即可满足条件：leftSum + num[i] = sum - leftSum => 2 * leftSum + num[i] = sum。
题目提到如果存在多个索引，则返回最左边那个，因此从左开始求和，而不是从右边。
*/
func pivotIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for index, num := range nums {
		// leftSum + num[i] = sum - leftSum.
		//左边和+当前元素=数组和-左边元素 即中心下标位置
		if leftSum*2+num == sum {
			return index
		}
		leftSum += num
	}
	return -1
}
