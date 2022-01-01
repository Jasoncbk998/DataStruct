/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/15 3:52 下午
 **/
package main

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
*/

func maxArea(height []int) int {
	max := 0
	for left, right := 0, len(height)-1; left < right; {
		//算面积
		area := Min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		//左边高,则移动右边,反之亦然
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
