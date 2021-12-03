/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 10:59 上午
 **/
package main

func main() {

}

//出让数组所有元素相等的最小操作次数
// 数组中所有元素都减少到数组中元素最小值所需的操作数
func minMoves(nums []int) (ans int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return
}
