/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/22 12:58 下午
 **/
package main

func main() {

}

/**
Input: 121
Output: true

Input: -121
Output: false

int%10 取个位
int/10 百位边十位  十位变个位
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	//将int -> 数组
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
