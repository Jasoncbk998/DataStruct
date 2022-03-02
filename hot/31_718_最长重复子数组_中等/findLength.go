/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/2 10:40 下午
 **/
package main

import "fmt"

/**
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
*/
// dp
/**
	1	2	3	2	1
3	0	0	1	0	0
2	0	1	0	1	0
1	1	0	0	0	1
4	0	0	0	0	0
7	0	0	0	0	0
*/
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	length := findLength([]int{5, 2, 3, 4, 5}, []int{1, 2, 3})
	fmt.Println(length)
}
