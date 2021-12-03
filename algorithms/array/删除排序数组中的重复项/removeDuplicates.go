/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 1:35 下午
 **/
package 删除排序数组中的重复项

/**

给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = 删除排序数组中的重复项(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
print(nums[i]);
}

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/

/**
快慢指针
通过两个指针去比对大小
*/
// 快慢指针
func RemoveDuplicates(nums []int) int {
	// arr长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	// 循环
	// 1 1 2 2 3 3
	//因为是有序且重复的数组,所以相邻两位进行比较,如果不等则快慢指针进行赋值替换
	for fast := 1; fast < n; fast++ {
		// 后位比较前位,比较相邻两位是否相等
		if nums[fast] != nums[fast-1] {
			//如果后位不等于前位,则把fast的值赋给slow
			nums[slow] = nums[fast]
			// slow 右移一位
			slow++
		}
	}
	return slow
}

func removeDuplicates1(nums []int) int {
	length := len(nums)
	slow := 1
	if length == 0 {
		return length
	}
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow

}
