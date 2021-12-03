/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 12:05 下午
 **/
package _2_removeElement

/**
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = 02_removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
	  print(nums[i]);
}

*/

/**
nums= 3 2 2 3 val = 3 去除3
*/

// 2 2 2 1  2
func RemoveElement(nums []int, val int) int {
	// 从索引位置0开始
	left := 0
	// 遍历数组
	for _, v := range nums {
		//进行比较,如果数组值!=目标值,则进行覆盖,并left++
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// nums=[3,2,2,3] val =3
// 2 nums[2,2]
func RemoveElement1(nums []int, val int) int {
	index := 0
	length := len(nums) - 1
	for i := 0; i <= length; i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index

}
