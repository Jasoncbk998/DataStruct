/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 12:05 下午
 **/
package _6_移除元素_简单

/**
nums= 3 2 2 3 val = 3 去除3
*/

func RemoveElement(nums []int, val int) int {
	//从最左边开始比较
	left := 0
	for _, v := range nums {
		//进行比较
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
