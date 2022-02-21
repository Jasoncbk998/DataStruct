/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2022/2/21 1:21 下午
 **/
package main

import "strconv"

const SEG_COUNT = 4

var (
	ans      []string
	segments []int
)

/**
这道题不会,挺难得,没看明白咋回事
*/
func restoreIpAddresses(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = []string{}
	afs(s, 0, 0)
	return nil
}

func afs(s string, segId, SegStart int) {
	if segId == SEG_COUNT {
		if SegStart == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != SEG_COUNT-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}
	if SegStart == len(s) {
		return
	}
	if s[SegStart] == '0' {

	}

}
