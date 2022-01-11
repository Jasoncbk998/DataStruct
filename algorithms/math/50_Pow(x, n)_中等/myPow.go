/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/20 12:02 下午
 **/
package main

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）。
输入：x = 2.00000, n = 10
输出：1024.00000
*/
func main() {
	println(myPow(2, 5))
}

//递归
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	//x的负一次方 就是这种情况
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
