/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/11/22 11:57 上午
 **/
package main

import "fmt"

var (
	letterMap = []string{
		" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"qprs",
		"tuv",
		"wxyz",
	}
	res   = []string{}
	final = 0
)

func main() {
	combinations := letterCombinations("23")
	fmt.Println(combinations)
}

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	d2c := [][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'},
		{'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'},
	}
	var bytes = make([]byte, len(digits))
	var ans = make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(digits) {
			ans = append(ans, string(bytes))
			return
		}
		for _, char := range d2c[digits[i]-'2'] {
			bytes[i] = char
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}

//非递归
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	//获取字母
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			//进行组合
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}
