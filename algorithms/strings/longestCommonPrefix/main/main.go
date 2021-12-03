/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/20 2:19 下午
 **/
package main

import (
	"DataStruct/algorithms/strings/longestCommonPrefix"
	"fmt"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix.LongestCommonPrefix(arr)
	fmt.Println(prefix)

}
