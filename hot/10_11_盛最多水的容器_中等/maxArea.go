/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/21 1:15 下午
 **/
package main

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
