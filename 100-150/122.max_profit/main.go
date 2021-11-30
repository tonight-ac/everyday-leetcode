package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}

// 寻找升序子序列
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i]-prices[i-1]
		}
	}
	return res
}
