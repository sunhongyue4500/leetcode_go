package main

import "fmt"

func main() {

	test2 := [][]int{
		[]int{2, 4, 1},
		[]int{3, 2, 6, 5, 0, 3},
		[]int{3, 2, 6, 5, 0, 3, 4, 10, 3, 1, 1, 2, 4, 19, 34, 23, 13, 25, 45, 100, 34, 56, 46, 89, 34},
	}
	k := 2
	for i := 0; i < len(test2); i++ {
		fmt.Println(maxProfit(k, test2[i]))
	}
	fmt.Println()
	for i := 0; i < len(test2); i++ {
		fmt.Println(maxProfit2(k, test2[i]))
	}
	fmt.Println()
	for i := 0; i < len(test2); i++ {
		fmt.Println(maxProfit3(k, test2[i]))
	}
	//fmt.Println(maxProfit(5, test2[2]))
}

// DP
// 注意：卖出才算一次交易完成
// 状态的定义：dp[i]，表示到第i天的最大利润
// dp[i][j]，表示第i天持有股票与否的最大利润
// dp[i][j][k]，表示第i天是否持有股票并且交易了k次的最大利润
// dp[i][k][0] 表示不持有股票， dp[i][k][1] 表示持有一只股票的最大利润
// 有多少个维度，一般就需要几个循环
// 遍历天数以及交易次数
// dp[i][k][0] = max{dp[i-1][k][0], dp[i-1][k-1][1] + price[i]}  不动或者卖出
// dp[i][k][1] = max{dp[i-1][k][0]-price[i], dp[i-1][k][1]}  不动或者买入
// ⚠️：分配内存的时候会OOM，比如有1000000000多天
// 如果k巨大无比，类似与可以进行无数次交易1000000000
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k < 1 {
		return 0
	}
	res := 0
	if k > len(prices) {
		for i := 1; i < len(prices); i++ {
			if prices[i]-prices[i-1] > 0 {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	// 初始化数据结构
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1) // 会有空间浪费
	}
	// 初始化
	dp[0][0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0][1] = max(dp[i-1][0][1], -prices[i]) // 持有之前的，或者持有今天的，看谁利润大
	}
	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}

	// 递推
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ { // 注意k的取值，最多可交易k次
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i]) // 不动或卖出
			dp[i][j][1] = max(dp[i-1][j][0]-prices[i], dp[i-1][j][1])   // 买入(不算交易)，或不动
			if dp[i][j][0] > res {
				res = dp[i][j][0]
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 如果k巨大无比，类似与可以进行无数次交易1000000000
// 状态压缩
func maxProfit2(k int, prices []int) int {
	if len(prices) == 0 || k < 1 {
		return 0
	}
	res := 0
	// 也就是说进行的交易次数，比天数还多，则特殊处理，防止k太大，导致分配超大数组，引起OOM
	if k > len(prices) {
		for i := 1; i < len(prices); i++ {
			if prices[i]-prices[i-1] > 0 {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	// DP
	// 初始化数据结构
	// 只需要相邻两天的状态
	dpLast := make([][2]int, k+1) // 前一天状态[0...k]
	dp := make([][2]int, k+1)     // 当前天状态
	// 初始化
	dpLast[0][1] = -prices[0]
	for i := 1; i <= k; i++ {
		dpLast[i][1] = max(dpLast[i-1][1], -prices[0]) // 持有之前的，或者持有今天的，看谁利润大
	}
	// 递推
	for i := 1; i < len(prices); i++ {
		dp[0][1] = max(dpLast[0][1], -prices[i])
		for j := 1; j <= k; j++ { // 注意k的取值，最多可交易k次
			dp[j][0] = max(dpLast[j][0], dpLast[j-1][1]+prices[i]) // 不动或卖出
			dp[j][1] = max(dpLast[j][0]-prices[i], dpLast[j][1])   // 买入(不算交易)，或不动
		}
		// 保存状态
		copy(dpLast, dp)
	}
	return dp[k][0]
}

// 别人的解法，只用一维数组
// 1.第k次buy就是 上一次卖出剩下的钱 - 本次购买需要的钱
// 2.第k次sell就是 第k次买入 + 本次卖出的钱

// buy[j] = max(buy[j], sell[j-1] - prices[i])
// sell[j] = max(sell[j], buy[j] + prices[i])
// 相比上面的解法更简洁
func maxProfit3(k int, prices []int) int {
	if k >= len(prices)/2 {
		ret := 0
		for i := 1; i < len(prices); i++ {
			if v := prices[i] - prices[i-1]; v > 0 {
				ret += v
			}
		}
		return ret
	}
	buy := make([]int, k+1)  // 第i天买入的最大利润
	sell := make([]int, k+1) // 第i天卖出的最大利润
	// 初始化买入
	for i := 0; i <= k; i++ {
		buy[i] = -2 << 31
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i]) //不动，前一天卖出后再买入
			sell[j] = max(sell[j], buy[j]+prices[i])  // 不动，上次买入后卖出
		}
	}
	return sell[k] // k次交易后的最大值
}
