/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/12/18 7:33 下午
 **/
package main

/**
输入: num1 = "2", num2 = "3"
输出: "6"
*/
/**
用数组模拟乘法。创建一个数组长度为 len(num1) + len(num2) 的数组用于存储乘积。
对于任意 0 ≤ i < len(num1)，0 ≤ j < len(num2)，num1[i] * num2[j] 的结果位于 tmp[i+j+1]，
如果 tmp[i+j+1]≥10，则将进位部分加到 tmp[i+j]。最后，将数组 tmp 转成字符串，如果最高位是 0 则舍弃最高位。
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	b1, b2, tmp := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(b2); j++ {
			tmp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}
