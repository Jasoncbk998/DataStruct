/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 12:05 下午
 **/
package main

/**
nums= 3 2 2 3 val = 3 去除3
*/
func RemoveElement(nums []int, val int) int {
	//标识符,从0开始对比
	flag := 0
	for _, v := range nums {
		if v != val {
			nums[flag] = v
			flag++
		}
	}
	return flag
}
