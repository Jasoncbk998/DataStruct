/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 4:28 下午
 **/
package romanToInt

/**
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做II，即为两个并列的 1。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。
这个特殊的规则只适用于以下六种情况：

I可以放在v (5) 和X(10) 的左边，来表示 4 和 9。
X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
C可以放在D(500) 和M(1000) 的左边，来表示400 和900。

输入: "III"
输出: 3

*/
//var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

var maps = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func RomanToInt(s string) int {
	n := len(s)
	ans := 0
	for i := range s {
		if i < n-1 && maps[s[i]] < maps[s[i+1]] {
			ans -= maps[s[i]]
		} else {
			ans += maps[s[i]]
		}
	}
	return ans
}
