/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/8 1:11 下午
 **/
package main

import (
	"fmt"
)

func main() {
	//最长公共子序列的长度是3
	ints1 := []int{1, 4, 3, 5, 9, 10}
	ints2 := []int{1, 4, 3, 9, 10}

	length := LCS3(ints1, ints2)
	fmt.Println(length)

}

// LCS3 2行的数组,继续退化
func LCS3(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([]int, len(nums1)+1)

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = Max(dp[j], dp[j-1])
			}
		}
	}
	return dp[len(nums2)]
}

// LCS2 退化二维数组,变成只有两行的,滚动数组
func LCS2(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		row := i % 2           //当前行
		prevRow := (i - 1) % 2 // 上一行
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[row][j] = dp[prevRow][j-1] + 1
			} else {
				dp[row][j] = Max(dp[prevRow][j], dp[row][j-1])
			}
		}
	}
	return dp[len(nums1)%2][len(nums2)]
}

//使用二维数组
func LCS1(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	// 动态创建二维数组
	//但是这个二维数组也很浪费,有很多数据都用不到
	// 其实是一个滚动数组,只用两行
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

// 改造递归,减少重复计算
func LCS(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	return lcs(nums1, nums2, len(nums1), len(nums2))
}

func lcs(nums1, nums2 []int, i, j int) int {
	// 递归基
	if i == 0 || j == 0 {
		return 0
	}
	if nums1[i-1] == nums2[j-1] {
		return lcs(nums1, nums2, i-1, j-1) + 1
	}
	return Max(lcs(nums1, nums2, i-1, j), lcs(nums1, nums2, i, j-1))

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
