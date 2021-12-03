/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/1 9:12 下午
 **/
package main

func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
func sortColors2(nums []int) {
	index := 0
	for i := range nums {
		if nums[i] == 0 {
			swap75(nums, index, i)
			index++
		}
	}
	for i := range nums {
		if nums[i] == 1 {
			swap75(nums, index, i)
			index++
		}
	}
}

func swap75(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
