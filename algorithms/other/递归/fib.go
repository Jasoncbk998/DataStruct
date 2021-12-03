/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/3 10:02 下午
 **/
package main

//最烂的递归
func fib_1(n int) int {
	if n <= 2 {
		return 1
	}
	return fib_1(n-1) + fib_1(n-2)
}

//加强版递归
func fib_2(n int) int {
	if n <= 2 {
		return 1
	}
	ints := make([]int, n+1)
	ints[1], ints[2] = 1, 1
	return fib_sum(n, ints)
}
func fib_sum(n int, ints []int) int {
	//如果ints[n]=0,那么就说明 ints[n]=0没被计算过,所以进行存储
	if ints[n] == 0 {
		//把递归中的中间结果值都存储到数组中
		ints[n] = fib_sum(n-1, ints) + fib_sum(n-2, ints)
	}
	return ints[n]
}

//去递归-进行调用
func fib_3(n int) int {
	if n <= 2 {
		return 1
	}
	ints := make([]int, n+1)
	ints[1], ints[2] = 1, 1
	for i := 3; i <= n; i++ {
		ints[i] = ints[i-1] + ints[i-2]
	}
	return ints[n]
}

//去递归-进行调用-用滚动数组进行优化
//因为每次相加都是用数组的两个元素,所以没必要存储那么多
func fib_4(n int) int {
	if n <= 2 {
		return 1
	}
	ints := make([]int, 2)
	ints[0], ints[1] = 1, 1
	for i := 3; i <= n; i++ {
		//因为取模效率很低,所以进行优化
		//ints[i%2] = ints[(i-1)%2] + ints[(i-2)%2]
		// 一定要注意,只有模2 才可以这么做
		ints[i&1] = ints[(i-1)&1] + ints[(i-2)&1]
	}
	//return ints[n%2]
	return ints[n&1]
}

//如果已经降低为滚动数组,只有两个元素,那么就可以退化为2个元素
func fib_5(n int) int {
	if n <= 2 {
		return 1
	}
	first, second := 1, 1
	for i := 3; i <= n; i++ {
		second = first + second
		first = second - first
	}
	return second
}
