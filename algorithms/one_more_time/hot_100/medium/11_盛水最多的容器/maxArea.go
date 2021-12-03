/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/25 3:02 下午
 **/
package main

/**
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/
func maxArea2(height []int) int {
	max_area := 0
	start := 0
	end := len(height) - 1

	for i := 0; ; i++ {
		if start == end {
			break
		}
		for ; start < end; start++ {
			if height[start] >= i {
				break
			}
		}
		for ; start < end; end-- {
			if height[end] >= i {
				break
			}
		}

		area := (end - start) * i
		if area > max_area {
			max_area = area
		}
	}

	return max_area
}
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	//扫描
	for start < end {
		width := end - start
		high := 0
		//通过比对start和end 确定high
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		//每次循环确定temp,并且会更新max
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

func test(height []int) int {
	max, start, end := 0, 0, len(height)-1

	for start < end {
		high := 0
		width := start - end
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			temp = max
		}

	}
	return max
}
