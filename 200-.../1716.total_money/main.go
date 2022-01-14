package main

import "fmt"

func main() {
	fmt.Println(totalMoney(4))
}

func totalMoney(n int) int {
	// 第n周 7*n+21 从0开始

	div, mod := n/7, n%7

	tot := 0

	if div > 0 {
		tot += 7 * (((1 + div) * div) / 2) + 21 * div
	}

	if mod > 0 {
		tot += (div + 1) * mod + (0 + mod - 1) * mod / 2
	}

	return tot
}