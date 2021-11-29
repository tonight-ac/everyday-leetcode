package main

import "fmt"

func main()  {
	fmt.Println(numTrees(4))
}
// 不需要像95题一样真的去建树
var dp = [20]int{1:1}
func numTrees(n int) int {
	if dp[n] != 0 { return dp[n] }

	for i := 2; i <= n; i++ {
		if dp[i] != 0 { continue }
		for j := 1; j <= i; j++ {
			l, r := 0, 0
			if j - 1 <= 0 { l = 1 } else { l = dp[j - 1] }
			if j + 1 > i { r = 1 } else { r = dp[i - j] }
			dp[i] += l * r
		}
	}

	return dp[n]
}
