/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 3:52 下午
 **/
package main

import "fmt"

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
*/
func main() {
	ints := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(ints))
}

//双指针,进行夹逼
func maxArea2(height []int) int {
	max := 0
	for left, right := 0, len(height)-1; left < right; {
		area := Min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		area := Min(height[left], height[right]) * (right - left)
		ans = Max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//
func maxArea4(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		//更新高,不断与宽相乘
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		//max进行更新
		if temp > max {
			max = temp
		}
	}
	return max
}
