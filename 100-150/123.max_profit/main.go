package main


func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(prices)) }

	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices); j++ {
			if j > i {
				dp[i][j] = oneMaxProfit(prices[i:j+1])
			}
		}
	}

	// 分为左右两份
	res := 0
	for i := 0; i < len(prices); i++ {
		if res < dp[0][i] + dp[i][len(prices)-1] {
			res = dp[0][i] + dp[i][len(prices)-1]
		}
	}

	return res
}

func oneMaxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			if res < prices[i] - min {
				res = prices[i] - min
			}
		} else {
			min = prices[i]
		}
	}
	return res
}

// 寻找能赚钱的交易，选最大两次
//func maxProfit(prices []int) int {
//	start := 0
//	wave := make([][]int, 0)
//
//	for i := 1; i < len(prices); i++ {
//		if prices[i] <= prices[i-1] {
//			if prices[i] != prices[i-1] {
//				wave = append(wave, []int{start, i-1})
//			}
//			start = i
//		}
//	}
//
//	res := 0
//	// 可操作波段数量小于2
//	if len(wave) <= 2 {
//		for i := 0; i < len(wave); i++ {
//			res += prices[wave[i][1]] - prices[wave[i][0]]
//		}
//	}
//
//	// 需要对wave搜索
//	return 0
//}