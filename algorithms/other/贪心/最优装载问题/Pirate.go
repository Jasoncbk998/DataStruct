/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/6 1:34 下午
 **/
package main

import (
	"fmt"
	"sort"
)

/**
贪心策略
*/
/**
装古董,怎么装能装最多
最优装载问题
船装在重量为W,每件古董重量为w,如何尽可能多的将古董装上船
每一次都选择重量最小的古董
*/

func main() {
	arr := []int{7, 1, 2, 4, 5, 6, 12, 64, 7, 3}
	Pirate(arr)
}

// Pirate 这个就是贪心的思想,每一步就是当前的最优选,局部最优选达到全局最优选
func Pirate(weights []int) {
	//先对数组进行排序
	sort.Ints(weights)
	capacity, weight, count := 30, 0, 0
	for i := 0; i < len(weights) && weight < capacity; i++ {
		newWeight := weight + weights[i]
		if newWeight <= capacity {
			weight = newWeight
			count++
			fmt.Println(weights[i])
		}
	}
	fmt.Println("一共选了", count, "个古董")
}
