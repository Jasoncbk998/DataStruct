/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 1:35 下午
 **/
package main

/**

给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = 80_删除有序数组中的重复项_简单(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
print(nums[i]);
}
输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/
// 快慢指针 没什么难度
func RemoveDuplicates(nums []int) int {
	// arr长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	//slow是进行换位的指针
	//fast是进行比较的指针
	slow := 1
	for fast := 1; fast < n; fast++ {
		// 1 1 2 2 3 4 5
		// 1 2 3 3 4
		//找不相等,覆盖相等
		//遇到相等的元素,则跳过,发现不相等,覆盖slow指针
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func test(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		// 1 2 2 3 3 4
		if arr[fast] != arr[fast-1] {
			arr[slow] = arr[fast]
			slow++
		}
	}
	return slow
}
