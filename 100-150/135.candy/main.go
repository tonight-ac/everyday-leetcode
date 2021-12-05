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

	res := 0
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
		res += candies[i]
	}

	return res + candies[len(candies)-1] + len(ratings)
}
