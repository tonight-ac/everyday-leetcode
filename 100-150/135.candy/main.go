package main

func candy(ratings []int) int {
	// 1 2 3 2 1
	// 0 1 2 1 0
	candies := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1] + 1)
		}
	}

	// 每人至少分1个
	res := len(ratings)
	for i := 0; i < len(candies); i++ {
		res += candies[i]
	}

	return res
}

func max(a, b int) int {
	if a > b { return a }
	return b
}