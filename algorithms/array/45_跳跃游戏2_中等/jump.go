/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/19 2:07 下午
 **/
package main

/**
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
意思是,在第一个元素,我们可以跳一步,步长是2,也可以跳两步,步长是1,跳到索引是2的元素
输入: nums = [2,3,1,1,4]
输出: 2
第一次:元素是2,最大步长是2,但是我们步长选择是1,跳到nums[1]=3,然后步长是3,直接跳到4 最末尾
*/

/**
要求找到最少跳跃次数，顺理成章的会想到用贪心算法解题。
扫描步数数组，维护当前能够到达最大下标的位置，记为能到达的最远边界，如果扫描过程中到达了最远边界，更新边界并将跳跃次数 + 1。
扫描数组的时候，其实不需要扫描最后一个元素，因为在跳到最后一个元素之前，
能到达的最远边界一定大于等于最后一个元素的位置，不然就跳不到最后一个元素，到达不了终点了；
如果遍历到最后一个元素，说明边界正好为最后一个位置，最终跳跃次数直接 + 1 即可,也不需要访问最后一个元素
*/
func main() {
	nums := []int{2, 3, 1, 2, 4, 2, 3}
	// 0 1
	println(jump(nums))

}
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
