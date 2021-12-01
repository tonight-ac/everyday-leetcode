package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			if  res < prices[i] - min {
				res = prices[i] - min
			}
		} else {
			min = prices[i]
		}
	}
	return res
}
