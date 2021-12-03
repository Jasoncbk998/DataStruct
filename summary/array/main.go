/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/18 2:34 下午
 **/
package array

import "fmt"

func main() {
	//索引 0,1,2,3,4,5,6,7,8 共9个索引,每个索引分别对应不同的value
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//索引,0,1,2,3,4  不包含5  前包后不包
	ints := arr[:5]
	for index, value := range ints {
		fmt.Println(index, value)
	}
	ints = arr[5:]
	fmt.Println()
	// 原数组索引 位置5开始到最后,包含5  前包后不包
	// 6,7,8,9
	for index, value := range ints {
		fmt.Println(index, value)
	}

	fmt.Println()
	//索引 索引 2,3,4,5
	//对应的值: 3,4,5,6
	//然后形成一个切片
	ints = arr[2:6]
	for index, value := range ints {
		fmt.Println(index, value)
	}

	test() // [0 0 0 0 0 0 0 0 0 0]

}

func test() {
	make_arr := make([]int, 10)
	fmt.Println(make_arr, &make_arr)

	new_arr := new([]int)
	fmt.Println(new_arr, &new_arr, *new_arr)

}
