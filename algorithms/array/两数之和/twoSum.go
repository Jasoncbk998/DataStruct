/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 12:41 下午
 **/
package 两数之和

/**
两数之和
给一个数组,和一个目标值,找到数组内两个值和为目标值得数
*/
func Sum2Number(arr []int, target int) []int {
	//构建map结构
	hashTable := map[int]int{}
	// 循环数组
	for i, x := range arr {
		// 利用map的k-v结构,进行比对,查询
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
