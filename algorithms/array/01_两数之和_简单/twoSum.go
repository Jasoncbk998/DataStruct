/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:41 下午
 **/
package main

import "fmt"

/**
01_两数之和_简单
给一个数组,和一个目标值,找到数组内两个值和为目标值得数
*/
func Sum2Number(arr []int, target int) []int {
	//构建map结构
	hashTable := map[int]int{}
	// 循环数组
	// idx,value
	for i, x := range arr {
		fmt.Println(i, "--", x)
		// 利用map的k-v结构,进行比对,查询
		//p是索引,key存索引,value存的是值
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	number := Sum2Number(ints, 3)
	fmt.Println(number)

}
