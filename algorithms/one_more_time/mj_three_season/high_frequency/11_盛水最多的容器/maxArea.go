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

//夹逼的思想
func maxArea2(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		//每一次找最小高度
		h := min(height[l], height[r])
		ans = max(ans, (r-l)*h)
		//不断移动
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func test(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		high := min(height[l], height[r])
		ans = max(ans, (r-l)*high)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
