package main

import "fmt"

//[3,3,5,0,0,3,1,4]
func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
}

// 寻找能赚钱的交易，选最大两次
func maxProfit(prices []int) int {
	dp1 := make([]int, len(prices))
	// 从左向右做多
	min, max := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
			max = prices[i]
		}
		if max < prices[i] { max = prices[i] }
		if max - min > 0 { dp1[i] = max - min }
	}

	dp2 := make([]int, len(prices))
	// 从右向左做空
	min, max = prices[len(prices)-1], prices[len(prices)-1]
	for i := len(prices)-2; i >= 0; i-- {
		if max < prices[i] {
			max = prices[i]
			min = prices[i]
		}
		if min > prices[i] { min = prices[i] }
		if max - min > 0 { dp2[i] = max - min }
	}

	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if res < dp1[i] + dp2[i+1] {
			res = dp1[i] + dp2[i+1]
		}
	}

	return res
}

// 超时了nmd
//func maxProfit(prices []int) int {
//	res := 0
//	for i := 0; i < len(prices); i++ {
//		left := oneMaxProfit(prices[:i])
//		right := oneMaxProfit(prices[i:])
//		if res < left + right {
//			res = left + right
//		}
//	}
//	return res
//}

//func oneMaxProfit(prices []int) int {
//	if len(prices) == 0 { return 0 }
//	res := 0
//	min := prices[0]
//	for i := 1; i < len(prices); i++ {
//		if prices[i] > min {
//			if res < prices[i] - min {
//				res = prices[i] - min
//			}
//		} else {
//			min = prices[i]
//		}
//	}
//	return res
//}

// 空间超了，搞一份不需要存储的
//func maxProfit(prices []int) int {
	//dp := make([][]int, len(prices))
	//for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(prices)) }

	//for i := 0; i < len(prices); i++ {
	//	for j := i+1; j < len(prices); j++ {
	//		dp[i][j] = oneMaxProfit(prices[i:j+1])
	//	}
	//}
	//
	//分为左右两份
	//res := 0
	//for i := 0; i < len(prices); i++ {
	//	if res < dp[0][i] + dp[i][len(prices)-1] {
	//		res = dp[0][i] + dp[i][len(prices)-1]
	//	}
	//}

	//return res
//}

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
