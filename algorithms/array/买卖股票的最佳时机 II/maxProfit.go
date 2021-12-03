/**
 * @Author liuxu22
 * @Description //TODO
 * @Date 2021/10/26 9:52 下午
 **/
package main

/**
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

输入: prices = [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

输入: prices = [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
//
//输入: prices = [7,1,5,3,6,4]
//输出: 7
func maxProfit4(prices []int) int {
	var sum int
	for i := 1; i < len(prices); i++ {
		//因为是卖股票,所以做差,大于0 就是赚
		// n天  -  n-1天
		//因为题目要求,多次买卖一直股票,只能持有一只股票,也就是说你在买之前,一定要卖掉手里的股票
		//,所以在一个for里边,通过累计sum进行最大利润统计
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

//动态规划
func maxProfit3(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

//贪心
func maxProfit2(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}
