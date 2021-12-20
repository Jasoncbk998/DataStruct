/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 12:05 下午
 **/
package _6_移除元素

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
