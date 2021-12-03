/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/6 10:35 下午
 **/
package main

import (
	"fmt"
	"math"
)

/**
贪心,找到的不是最优解
动态规划可以找到最优解
*/

/**
零钱问题
假设有25分,10分,5分,1分的硬币,心在要找给客户41分的零钱,如何做到硬币个数最少
思路就是优先选择面值最大的硬币

最优的是用动态规划的思想去做
*/

func main() {
	i := coinChange([]int{1, 5, 20, 25}, 4) // 3枚硬币
	fmt.Println(i)
}

func coinChange(coins []int, amount int) int {
	var arr = make([]int, amount+1)
	arr[0] = 0
	for i := 1; i <= amount; i++ {
		arr[i] = 2147483647
		var cur = 0
		for _, val := range coins {
			cur = i - val
			if cur >= 0 {
				arr[i] = Min(arr[i], arr[cur]+1)
			}
		}
	}
	if arr[amount] == 2147483647 {
		return -1
	}

	return arr[amount]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coins5(amount int, coins []int) int {
	if amount < 1 || len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := math.MaxInt64
		for _, face := range coins {
			if i < face || dp[i-face] < 0 {
				continue
			}
			if dp[i-face] >= min {
				continue
			}
			min = Min(dp[i-face], min)
		}
		if min == math.MaxInt64 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}
	return dp[amount]
}
func coins_final(n int, faces []int) int {
	if n < 1 || len(faces) == 0 {
		return -1
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		for _, face := range faces {
			if i < face {
				continue
			}
			v := dp[i-face]
			if v < 0 || v >= min {
				continue
			}
			min = v
		}
		if min == math.MaxInt64 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}
	return dp[n]
}
func coins4(n int8) int8 {
	if n < 1 {
		return -1
	}
	dp := make([]int8, n+1)
	//凑够i分时最后选择的硬币面值
	face := make([]int8, len(dp))
	for i := 1; i <= int(n); i++ {
		mins := math.MaxInt8
		if i >= 1 && dp[i-1] < int8(mins) {
			mins = int(dp[i-1])
			face[i] = 1
		}
		if i >= 5 && dp[i-5] < int8(mins) {
			mins = int(dp[i-5])
			face[i] = 5
		}
		if i >= 20 && dp[i-20] < int8(mins) {
			mins = int(dp[i-20])
			face[i] = 20
		}
		if i >= 25 && dp[i-25] < int8(mins) {
			mins = int(dp[i-25])
			face[i] = 25
		}
		dp[i] = int8(mins + 1)
	}

	print(face, int(n))

	return dp[n]
}
func print(face []int8, n int) {
	for n > 0 {
		fmt.Println(face[n])
		n -= int(face[n])
	}
}

//递推,自底向上
func coins3(n int8) int8 {
	if n < 1 {
		return -1
	}
	dp := make([]int8, n+1)
	for i := 1; i <= int(n); i++ {
		mins := dp[i-1]
		if i >= 5 {
			mins = min(dp[i-5], mins)
		}
		if i >= 20 {
			mins = min(dp[i-20], mins)
		}
		if i >= 25 {
			mins = min(dp[i-25], mins)
		}
		dp[i] = mins + 1
	}
	return dp[n]
}

// 记忆化搜索,每个n只记录一次,避免重复计算
func coins2(n int8) int8 {
	if n < 1 {
		return -1
	}
	dp := make([]int8, n+1)
	faces := []int8{1, 5, 20, 25}
	for _, face := range faces {
		if n < face {
			break
		}
		dp[face] = 1
	}
	return Coins2(n, dp)
}

func Coins2(n int8, dp []int8) int8 {
	if n < 1 {
		return math.MaxInt8
	}
	if dp[n] == 0 {
		min1 := min(Coins2(n-25, dp), Coins2(n-20, dp))
		min2 := min(Coins2(n-5, dp), Coins2(n-1, dp))
		dp[n] = min(min1, min2) + 1
	}
	return dp[n]
}

//暴力递归,重叠子问题
//这种写法,涉及到了大量重复计算,具体可以参照递归调用
//所以进行优化升级改造
func coins(n int8) int8 {
	if n < 1 {
		// 如果n<1,那么返回int的最大值,这样相比较的另一个值就一定是最小值
		return math.MaxInt8
	}
	if n == 25 || n == 20 || n == 5 || n == 1 {
		return 1
	}
	// 凑够41,面值25分的硬币,最少需要几个
	min1 := min(coins(n-25), coins(n-20))
	min2 := min(coins(n-5), coins(n-1))
	return min(min1, min2) + 1

}

func min(a, b int8) int8 {
	if a <= b {
		return a
	}
	return b
}
