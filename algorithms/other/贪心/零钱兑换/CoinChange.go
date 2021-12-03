/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/6 1:50 下午
 **/
package main

import (
	"fmt"
	"sort"
)

/**
零钱问题
假设有25分,10分,5分,1分的硬币,心在要找给客户41分的零钱,如何做到硬币个数最少
思路就是优先选择面值最大的硬币
*/

func main() {
	faces := []int{25, 20, 5, 1}
	CoinChange(41, faces)
}

//最优解,一种硬币,最多取几次
//比如面值25的硬币,最多取几次,取完后,在去取20面值的硬币,计算最多取几次
func CoinChange(money int, faces []int) {
	sort.Ints(faces)
	fmt.Println(faces)
	coins, idx := 0, len(faces)-1
	for idx >= 0 {
		for money >= faces[idx] {
			fmt.Println(faces[idx])
			//选择硬币
			money -= faces[idx]
			coins++
		}
		//选择面额合适的硬币
		idx--
	}
	fmt.Println(coins)
}

func CoinChange2(money int, faces []int) {
	sort.Ints(faces)
	coins, i := 0, 0
	for i < len(faces) {
		if money < faces[i] {
			i++
			continue
		}
		fmt.Println(faces[i])
		money -= faces[i]
		coins++
	}
	fmt.Println(coins)
}

func CoinChange3(money int, faces []int) {
	sort.Ints(faces)
	coins := 0
	for i := len(faces) - 1; i >= 0; i-- {
		if money < faces[i] {
			continue
		}
		fmt.Println(faces[i])
		money -= faces[i]
		coins++
		i = len(faces)
	}
	fmt.Println(coins)
}
