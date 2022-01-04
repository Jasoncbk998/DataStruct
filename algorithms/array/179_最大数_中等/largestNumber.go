/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/10/20 12:20 下午
 **/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	str := largestNumber(nums)
	fmt.Println(str)
}

/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
输入：nums = [10,2]
输出："210"
输入：nums = [3,30,34,5,9]
输出："9534330"
*/
func largestNumber(nums []int) string {
	//自己定义排序接口
	//自定义排序条件
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		//这个数学公式有点难搞
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

/**
在比较 2 个字符串大小的时候，不单纯的只用字符串顺序进行比较，还加入一个顺序。
aStr := a + b
bStr := b + a
通过比较 aStr 和 bStr 的大小来得出是 a 大还是 b 大。
举个例子，还是 “3” 和 “30” 的例子，比较这 2 个字符串的大小。
aStr := "3" + "30" = "330"
bStr := "30" + "3" = "303"
通过互相补齐位数之后再进行比较，就没有问题了。很显然这里 “3” 比 “30” 要大。
*/
func largestNumber1(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numStrs := toStringArray(nums)
	quickSortString(numStrs, 0, len(numStrs)-1)
	res := ""
	for _, str := range numStrs {
		if res == "0" && str == "0" {
			continue
		}
		res = res + str
	}
	return res
}

func toStringArray(nums []int) []string {
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	return strs
}
func partitionString(a []string, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		ajStr := a[j] + pivot
		pivotStr := pivot + a[j]
		if ajStr > pivotStr { // 这里的判断条件是关键
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSortString(a []string, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionString(a, lo, hi)
	quickSortString(a, lo, p-1)
	quickSortString(a, p+1, hi)
}
