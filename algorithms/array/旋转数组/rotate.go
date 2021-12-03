/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/27 10:23 上午
 **/
package main

import "fmt"

/**
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
*/
func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(ints, k)
	fmt.Println(ints)
}

//执行超时,太慢了
func rotate(nums []int, k int) {
	length := len(nums)
	temp := 0
	for i := 0; i < k; i++ {
		temp = nums[length-1]
		for j := length - 1; j < length && j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}

func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	// 函数 copy 在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 指向同⼀底层数组。直接对应位置覆盖。
	copy(nums, newNums)
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse1(nums, 0, len(nums)-1)
	reverse1(nums, 0, k-1)
	reverse1(nums, k, len(nums)-1)

}
func reverse1(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
