/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 2:19 下午
 **/
package main

import (
	"DataStruct/algorithms/strings/最长公共前缀"
	"fmt"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	prefix := 最长公共前缀.LongestCommonPrefix(arr)
	fmt.Println(prefix)

}
