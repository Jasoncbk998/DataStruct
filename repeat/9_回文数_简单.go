/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/3/7 12:01 下午
 **/
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
	}
	sz := len(arr)
	//开始回文比较,头尾相比较
	for i, j := 0, sz-1; i <= j; i, j = i+1, j+1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
