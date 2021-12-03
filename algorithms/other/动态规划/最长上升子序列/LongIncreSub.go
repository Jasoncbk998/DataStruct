/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/7 6:37 下午
 **/
package main

import "fmt"

func main() {
	//不用连续
	//长度是4 ,2,4,7,101  一共4个
	num := []int{10, 2, 2, 5, 1, 7, 101, 18}
	sub := lengthOfLIS3(num)
	fmt.Println(sub)
}

func lengthOfLIS4(nums []int) int {
	if len(nums) == 0 || nums != nil {
		return 0
	}
	length := 0
	top := make([]int, len(nums))
	for _, num := range nums {
		begin := 0
		end := length
		for begin < end {
			mid := (begin + end) >> 1
			if num <= top[mid] {
				end = mid
			} else {
				begin = mid + 1
			}
		}
		//覆盖牌顶
		top[begin] = num
		//检查是否需要新建一个牌堆
		if begin == length {
			length++
		}
	}
	return length

}

//二分搜索的方法去写
func lengthOfLIS3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0

	top := make([]int, len(nums))
	for _, num := range nums {
		j := 0
		for j < length {
			if top[j] >= num {
				//覆盖牌顶
				top[j] = num
				break
			}
			j++
		}
		//新建一个牌堆
		if j == length {
			length++
			top[j] = num
		}
	}
	return length

}

//求出最大上升子序列是谁就行了,然后计算一下长度
//我们可以看一下,以每个元素结尾的子序列是不是上升的,
// 比如dp(0) 是以10结尾的子序列,只有一个元素
//dp(1) 是10,2 2结尾的子序列,不是上升
//dp(2) 是10,2,2 这个时候也没有
//dp(3) 是5结尾,那么是2,5 上升 是2位
//以此类推

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			dp[i] = Max(dp[i], dp[j]+1)
		}
		max = Max(dp[i], max)
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//普通解法
func lengthOfLIS(nums []int) int {
	var dp []int
	first := 1
	for _, num := range nums {
		if first == 1 {
			first = 0
			dp = append(dp, num)
			continue
		}
		dpLength := len(dp)
		if num > dp[dpLength-1] {
			dp = append(dp, num)
		} else {
			left := 0
			right := dpLength
			pos := 0
			for left <= right {
				mid := (left + right) / 2
				if dp[mid] < num {
					pos = mid + 1
					left = pos
				} else {
					right = mid - 1
				}
			}
			dp[pos] = num
		}
	}
	return len(dp)
}
