/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/7 5:57 下午
 **/
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	array := MaxSubArray_test(nums)
	fmt.Println(array)
}

//因为dp如果用成数组,我们每次都是比较上一个元素,这样就会很浪费
//所以退化数组为一个值就行了
func maxSubArray1(num []int) int {
	if len(num) == 0 {
		return -1
	}
	dp := num[0]
	max := dp
	for i := 1; i < len(num); i++ {
		if dp <= 0 {
			dp = num[i]
		} else {
			dp = dp + num[i]
		}
		max = Max(dp, max)
	}
	return max
}

//使用动态规划去做
//分别计算,以-2结尾的最大连续子序列的和是-2,就是dp(0)=-2
// 以1结尾的最大连续子序列的和是dp(1)=1 ,就是只有一个1
//一定要注意,以谁结尾的最大连续子序列,必须包含这个元素,必须是他结尾
//以-3结尾的最大连续子序列的和dp(2)=-2= 1-3=-2
func maxSubArray(num []int) int {
	if len(num) == 0 {
		return -1
	}
	dp := make([]int, len(num))
	dp[0] = num[0]
	max := dp[0]
	for i := 1; i < len(num); i++ {
		//如果子序列小于等于0则忽略
		if dp[i-1] <= 0 {
			dp[i] = num[i]
		} else {
			//dp[i-1]>0,则相加,形成新的子序列
			dp[i] = dp[i-1] + num[i]
		}
		max = Max(dp[i], max)
		fmt.Println("dp[", i, "]:", dp[i])
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
