/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/20 12:02 下午
 **/
package main

/**
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

//位运算
func myPow1(x float64, n int) float64 {
	res, isNeg := float64(1), false
	if n < 0 {
		n, isNeg = -n, true
	}

	for num, multi := n, x; num != 0; num, multi = num>>1, multi*multi {
		if num&1 != 0 {
			res *= multi
		}
	}

	if isNeg {
		res = 1 / res
	}
	return res
}
