/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/18 9:16 下午
 **/
package main

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
将数学问题,用代码呈现出来,这个思维的转变,需要锻炼
*/
func maxArea(height []int) int {
	max := 0
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		area = Min(height[right], height[left]) * (right - left)
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
