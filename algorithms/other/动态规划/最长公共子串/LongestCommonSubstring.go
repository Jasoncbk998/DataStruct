/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/8 7:15 下午
 **/
package main

import "fmt"

func main() {
	str1 := "abcba"
	str2 := "abcd"
	ssubstring := LCSsubstring1(str1, str2)
	fmt.Println(ssubstring)

}

func LCSsubstring1(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	char1 := []rune(str1)
	if len(char1) == 0 {
		return 0
	}
	char2 := []rune(str2)
	if len(char2) == 0 {
		return 0
	}
	colsChars := char1
	rowsChars := char2

	if len(char1) < len(char2) {
		colsChars = char1
		rowsChars = char2
	}

	dp := make([]int, len(colsChars)+1)

	max := 0
	for row := 1; row < len(rowsChars); row++ {
		cur := 0
		for col := 1; col < len(colsChars); col++ {
			leftTop := cur
			cur = dp[col]
			if char1[row-1] != char2[col-1] {
				dp[col] = 0
			} else {
				dp[col] = leftTop + 1
				max = Max(dp[col], max)
			}
		}
	}
	return max
}

func LCSsubstring(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	char1 := []rune(str1)
	if len(char1) == 0 {
		return 0
	}
	char2 := []rune(str2)
	if len(char2) == 0 {
		return 0
	}
	dp := make([][]int, Min(len(char1), len(char2)))

	for i := 0; i < len(char1)+1; i++ {
		dp[i] = make([]int, len(char2)+1)
	}
	max := 0
	for i := 1; i < len(char1); i++ {
		for j := 1; j < len(char2); j++ {
			if char1[i-1] != char2[j-1] {
				continue
			}
			//dp[i-1]=dp[j-1]时,退化到上一个状态进行求解
			dp[i][j] = dp[i-1][j-1] + 1
			//更新max
			max = Max(dp[i][j], max)
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
