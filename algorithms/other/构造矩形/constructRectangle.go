/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/25 10:04 上午
 **/
package main

import "math"

func main() {

}

/**
构造矩形
输入: 4
输出: [2, 2]
解释: 目标面积是 4， 所有可能的构造方案有 [1,4], [2,2], [4,1]。
但是根据要求2，[1,4] 不符合要求; 根据要求3，[2,2] 比 [4,1] 更能符合要求. 所以输出长度 L 为 2， 宽度 W 为 2。
*/
func constructRectangle(area int) []int {
	res := []int{}
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			res = []int{area / i, i}
		}
	}
	return res
}
func constructRectangle1(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}
